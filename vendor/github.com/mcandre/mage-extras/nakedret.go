package mageextras

import (
	"os"
	"os/exec"
)

// Nakedret runs nakedret.
func Nakedret(args ...string) error {
	cmd := exec.Command("nakedret")
	cmd.Args = append(cmd.Args, args...)
	cmd.Args = append(cmd.Args, "./...")
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
