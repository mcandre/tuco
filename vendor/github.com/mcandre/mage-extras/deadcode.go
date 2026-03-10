package mageextras

import (
	"os"
	"os/exec"
)

// Deadcode runs deadcode.
func Deadcode(args ...string) error {
	cmd := exec.Command("deadcode")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
