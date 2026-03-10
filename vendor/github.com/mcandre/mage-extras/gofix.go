package mageextras

import (
	"os"
	"os/exec"
)

// GoFix runs go fix.
func GoFix(args ...string) error {
	cmd := exec.Command("go", "fix")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
