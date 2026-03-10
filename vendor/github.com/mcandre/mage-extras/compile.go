package mageextras

import (
	"os"
	"os/exec"
)

// Compile runs go build recursively.
func Compile(args ...string) error {
	cmd := exec.Command("go", "build")
	cmd.Args = append(cmd.Args, "build")
	cmd.Args = append(cmd.Args, args...)
	cmd.Args = append(cmd.Args, AllPackagesPath)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
