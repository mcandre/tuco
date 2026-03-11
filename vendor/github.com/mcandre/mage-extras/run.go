package mageextras

import (
	"errors"
	"os"
	"os/exec"
)

// Run executes commands with practical defaults.
//
// * Inherit parent environment variables
// * Inherit parent I/O handles
func Run(program string, args ...string) error {
	if len(program) == 0 {
		return errors.New("blank program name")
	}

	cmd := exec.Command(program)
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
