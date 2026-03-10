package mageextras

import (
	"os"
	"os/exec"
)

// Govulncheck runs govulncheck.
func Govulncheck(args ...string) error {
	cmd := exec.Command("govulncheck")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
