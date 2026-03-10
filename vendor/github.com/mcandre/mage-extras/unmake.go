package mageextras

import (
	"os"
	"os/exec"
)

// Unmake runs unmake.
func Unmake(args ...string) error {
	cmd := exec.Command("unmake")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
