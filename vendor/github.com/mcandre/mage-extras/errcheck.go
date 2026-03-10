package mageextras

import (
	"os"
	"os/exec"
)

// Errcheck runs errcheck.
func Errcheck(args ...string) error {
	cmd := exec.Command("errcheck")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
