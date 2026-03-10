package mageextras

import (
	"os"
	"os/exec"
)

// DockerScout executes a docker security audit.
func DockerScout(args ...string) error {
	cmd := exec.Command("docker")
	cmd.Args = append(cmd.Args, "scout")
	cmd.Args = append(cmd.Args, "cves")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
