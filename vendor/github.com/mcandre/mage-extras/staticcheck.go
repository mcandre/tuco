package mageextras

import (
	"os"
	"os/exec"
)

// Staticcheck runs staticcheck.
func Staticcheck(args ...string) error {
	cmd := exec.Command("staticcheck")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
