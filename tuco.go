// Package tuco implements primitives for organizing Go crosscompilation.
package tuco

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/magefile/mage/sh"
	"gopkg.in/yaml.v3"
)

// ConfigurationFilename denotes the location of the tuco configuration file,
// relative to the current working directory.
const ConfigurationFilename = "tuco.yaml"

// DefaultArtifacts denotes the default location of the artifact directory tree.
const DefaultArtifacts = "bin"

// DefaultJobs denotes the default number of goroutines.
const DefaultJobs uint = 4

// DefaultArchiveExcludes generates file path patterns to strip from binary archives.
var DefaultArchiveExcludes = sync.OnceValue(func() []string {
	return []string{
		".DS_Store",
		"Thumbs.db",
	}
})

// Port models a basic targetable execution configuration.
type Port struct {
	// Os denotes a high-level environment.
	//
	// Example: "linux"
	Os string

	// Arch denotes a low-level environment.
	//
	// Example: "amd64"
	Arch string
}

// ParsePort constructs a port from a string.
// https://go.dev/wiki/PortingPolicy
func ParsePort(s string) (*Port, error) {
	parts := strings.Split(s, "/")

	if len(parts) < 2 {
		return nil, fmt.Errorf("cannot parse port metadata: %v", s)
	}

	return &Port{Os: parts[0], Arch: parts[1]}, nil
}

// String renders a port.
func (o Port) String() string {
	return fmt.Sprintf("%s/%s", o.Os, o.Arch)
}

// MarshalYAML encodes YAML.
func (o Port) MarshalYAML() (interface{}, error) {
	aux := o.String()
	return yaml.Marshal(aux)
}

// UnmarshalYAML decodes YAML.
func (o *Port) UnmarshalYAML(value *yaml.Node) error {
	var aux string

	if err := value.Decode(&aux); err != nil {
		return err
	}

	p, err := ParsePort(aux)

	if err != nil {
		return err
	}

	*o = *p
	return nil
}

// Tuco models a crossport build setup.
type Tuco struct {
	// Debug enables additional logging.
	Debug bool `yaml:"debug,omitempty"`

	// Artifacts denotes the location of the toplevel artifacts directory (default: `DefaultArtifacts`).
	Artifacts string `yaml:"artifacts,omitempty"`

	// Banner denotes a software application identifier (required, nonblank).
	Banner string `yaml:"banner"`

	// Jobs limits the number of goroutines (default: `DefaultJobs`).
	Jobs uint `yaml:"jobs,omitempty"`

	// GoArgs collects additional `go build`... CLI flags
	GoArgs []string `yaml:"go_args,omitempty"`

	// Os enables matching Go ports by GOOS.
	Os []string `yaml:"os,omitempty"`

	// Arch enables matching Go ports by GOARCH.
	Arch []string `yaml:"arch,omitempty"`

	// PortExcludes skips named Go ports.
	PortExcludes []string `yaml:"port_excludes,omitempty"`

	// ArchiveExcludes skips matching file paths when archiving binaries.
	//
	// Glob syntax
	// https://pkg.go.dev/path/filepath#Match
	ArchiveExcludes []string `yaml:"archive_excludes,omitempty"`

	// ports caches target ports.
	ports []Port `yaml:"-"`

	// tarballRoot caches the root directory for binary archives.
	tarballRoot string `yaml:"-"`

	// maxPortLen caches the length of the longest "<os>/<arch>" combination
	maxPortLen int `yaml:"-"`
}

// UpdateTarballRoot calculates binary archive root directories.
func (o *Tuco) UpdateTarballRoot() {
	o.tarballRoot = path.Join(o.Artifacts, fmt.Sprintf("%s-ports", o.Banner))
}

// UpdateMaxPortLen calculates maxPortLen.
func (o *Tuco) UpdateMaxPortLen() {
	var maxPortLen int

	for _, port := range o.ports {
		portLen := len(port.String())

		if portLen > maxPortLen {
			maxPortLen = portLen
		}
	}

	o.maxPortLen = maxPortLen
}

// NewTuco constructs a default Tuco.
func NewTuco() Tuco {
	var tc Tuco
	tc.Artifacts = DefaultArtifacts
	tc.Jobs = DefaultJobs
	tc.ArchiveExcludes = DefaultArchiveExcludes()
	return tc
}

// Load constructs a Tuco from `ConfigurationFilename`.
func Load() (*Tuco, error) {
	tucoYAMLBytes, err := os.ReadFile(ConfigurationFilename)

	if err != nil {
		return nil, err
	}

	tc := NewTuco()
	decoder := yaml.NewDecoder(bytes.NewReader(tucoYAMLBytes))
	decoder.KnownFields(true)

	if err := decoder.Decode(&tc); err != nil {
		return nil, err
	}

	tc.UpdateTarballRoot()

	if err := tc.UpdatePortCache(); err != nil {
		return nil, err
	}

	tc.UpdateMaxPortLen()
	return &tc, nil
}

// Prefix generates log namespaces.
func (o Tuco) Prefix(port Port) string {
	portPadded := fmt.Sprintf("%-*s", o.maxPortLen, port)
	return fmt.Sprintf("[ %s ] ", portPadded)
}

// ExecutableOutputPath generates file paths for executables.
func (o Tuco) ExecutableOutputPath(port Port) string {
	return path.Join(o.Artifacts, o.Banner, port.String())
}

// Clean removes artifacts.
func (o Tuco) Clean() error {
	return sh.Rm(o.Artifacts)
}

// Archive compresses Go applications in conventional UNIX TGZ format.
func (o Tuco) Archive(port Port) error {
	prefix := o.Prefix(port)

	log.Printf("%sarchiving\n", prefix)

	tarballRoot := o.tarballRoot
	tarballBasename := fmt.Sprintf("%s-%s-%s.tgz", o.Banner, port.Os, port.Arch)
	tarball := path.Join(tarballRoot, tarballBasename)
	tarballFile, err := os.Create(tarball)

	if err != nil {
		return err
	}

	defer func() {
		if err2 := tarballFile.Close(); err2 != nil {
			log.Println(err2)
		}
	}()

	gw := gzip.NewWriter(tarballFile)
	defer func() {
		if err2 := gw.Close(); err2 != nil {
			log.Println(err2)
		}
	}()
	tw := tar.NewWriter(gw)
	defer func() {
		if err2 := tw.Close(); err2 != nil {
			log.Println(err2)
		}
	}()

	outputPth := o.ExecutableOutputPath(port)
	entries, err := os.ReadDir(outputPth)

	if err != nil {
		return err
	}

	for _, entry := range entries {
		info, err2 := entry.Info()

		if err2 != nil {
			return err2
		}

		header, err2 := tar.FileInfoHeader(info, "")

		if err2 != nil {
			return err2
		}

		basename := entry.Name()
		sourcePth := path.Join(outputPth, basename)

		var isExcluded bool

		for _, exclusion := range o.ArchiveExcludes {
			m, err := filepath.Match(exclusion, basename)

			if err != nil {
				return err
			}

			if m {
				isExcluded = true
				break
			}
		}

		if isExcluded {
			continue
		}

		if !info.Mode().IsRegular() {
			return fmt.Errorf("unsupported file system file type for path: %s", sourcePth)
		}

		// UNIX executables. Binaries, shell style scripts, etc.
		mode := int64(0755)

		// Windows executables, general purpose scripts, etc.
		if strings.Contains(basename, ".") {
			mode = int64(0644)
		}

		header.Mode = mode

		if err3 := tw.WriteHeader(header); err3 != nil {
			return err3
		}

		f, err2 := os.Open(sourcePth)

		if err2 != nil {
			return err2
		}

		defer func() {
			if err3 := f.Close(); err3 != nil {
				log.Println(err3)
			}
		}()

		if _, err3 := io.Copy(tw, f); err3 != nil {
			return err3
		}
	}

	return nil
}

// prefixStream disambiguates concurrent child process output streams.
func prefixStream(prefix string, wg *sync.WaitGroup, r io.Reader) {
	defer wg.Done()

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		log.Printf("%s%s\n", prefix, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		if !errors.Is(err, os.ErrClosed) {
			log.Printf("%serror: %v\n", prefix, err)
		}
	}
}

// Build generates binaries for the given port.
func (o Tuco) Build(port Port) error {
	prefix := o.Prefix(port)

	log.Printf("%scompiling\n", prefix)

	if len(o.Artifacts) == 0 {
		return errors.New("blank artifacts")
	}

	if len(o.Banner) == 0 {
		return errors.New("blank banner")
	}

	outputPth := o.ExecutableOutputPath(port)

	if err := os.MkdirAll(outputPth, 0755); err != nil {
		return err
	}

	allPackagesPath := fmt.Sprintf(".%c...", os.PathSeparator)

	cmd := exec.Command("go")
	cmd.Args = []string{"go", "build", "-o", outputPth}
	cmd.Args = append(cmd.Args, o.GoArgs...)
	cmd.Args = append(cmd.Args, allPackagesPath)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%s", port.Os))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOARCH=%s", port.Arch))

	stderrReader, err := cmd.StderrPipe()

	if err != nil {
		return err
	}

	stdoutReader, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	if o.Debug {
		log.Printf("%scommand: %s\n", prefix, cmd)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go prefixStream(prefix, &wg, stderrReader)
	go prefixStream(prefix, &wg, stdoutReader)

	if err2 := cmd.Start(); err2 != nil {
		return err2
	}

	err = cmd.Wait()
	wg.Wait()

	if err != nil {
		return fmt.Errorf("%serror: %v", prefix, err)
	}

	return o.Archive(port)
}

func (o *Tuco) UpdatePortCache() error {
	cmd := exec.Command("go", "tool", "dist", "list")
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr

	stdout, err := cmd.Output()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(bytes.NewReader(stdout))
	availableGoosSet := make(map[string]struct{})
	availableGoarchSet := make(map[string]struct{})
	availablePortsSet := make(map[Port]struct{})

	for scanner.Scan() {
		line := scanner.Text()
		portP, err2 := ParsePort(line)

		if err2 != nil {
			return err2
		}

		port := *portP
		availableGoosSet[port.Os] = struct{}{}
		availableGoarchSet[port.Arch] = struct{}{}
		availablePortsSet[port] = struct{}{}
	}

	goosSet := make(map[string]struct{})

	for _, goos := range o.Os {
		if _, ok := availableGoosSet[goos]; !ok {
			return fmt.Errorf("invalid goos: %s", goos)
		}

		goosSet[goos] = struct{}{}
	}

	goarchSet := make(map[string]struct{})

	for _, goarch := range o.Arch {
		if _, ok := availableGoarchSet[goarch]; !ok {
			return fmt.Errorf("invalid goarch: %s", goarch)
		}

		goarchSet[goarch] = struct{}{}
	}

	excludeSet := make(map[string]struct{})

	for _, exclusion := range o.PortExcludes {
		excludeSet[exclusion] = struct{}{}
	}

	var enabledPorts []Port

	for port := range availablePortsSet {
		if _, ok := goosSet[port.Os]; !ok {
			continue
		}

		if _, ok := goarchSet[port.Arch]; !ok {
			continue
		}

		if _, ok := excludeSet[port.String()]; ok {
			continue
		}

		enabledPorts = append(enabledPorts, port)
	}

	sort.Slice(enabledPorts, func(i, j int) bool {
		if enabledPorts[i].Os == enabledPorts[j].Os {
			return enabledPorts[i].Arch < enabledPorts[j].Arch
		}
		return enabledPorts[i].Os < enabledPorts[j].Os
	})
	o.ports = enabledPorts
	return nil
}

// Run crosscompiles and archives Go applications.
func (o *Tuco) Run() []error {
	goos := o.Os
	goosLen := len(goos)

	if goosLen == 0 {
		log.Println("warning: empty os")
		return nil
	}

	goarch := o.Arch
	goarchLen := len(goarch)

	if goarchLen == 0 {
		log.Println("warning: empty arch")
		return nil
	}

	tarballRoot := o.tarballRoot
	var errs []error

	if err := os.MkdirAll(tarballRoot, 0755); err != nil {
		errs = append(errs, err)
		return errs
	}

	ports := o.ports
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(ports))
	buildJobsCh := make(chan Port)

	for w := uint(1); w <= o.Jobs; w++ {
		go func(wg *sync.WaitGroup, m *sync.Mutex, errs *[]error) {
			for port := range buildJobsCh {
				if err := o.Build(port); err != nil {
					m.Lock()
					*errs = append(*errs, err)
					m.Unlock()
				}

				wg.Done()
			}
		}(&wg, &m, &errs)
	}

	for _, port := range ports {
		buildJobsCh <- port
	}

	wg.Wait()

	if len(errs) != 0 {
		return errs
	}

	log.Printf("binaries archived: %s\n", tarballRoot)

	return nil
}
