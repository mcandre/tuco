package mageextras

import (
	"os"
	"os/exec"
)

// Tuggy runs tuggy.
func Tuggy(args ...string) error {
	cmd := exec.Command("tuggy")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
