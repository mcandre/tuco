package mageextras

import (
	"os"
	"os/exec"
)

// Tuco runs tuco.
func Tuco(args ...string) error {
	cmd := exec.Command("tuco")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
