package mageextras

import (
	"os"
	"os/exec"
)

// Chandler runs chandler.
func Chandler(args ...string) error {
	cmd := exec.Command("chandler")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
