package mageextras

import (
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

// GoImports runs goimports.
func GoImports(args ...string) error {
	mg.Deps(CollectGoFiles)

	for pth := range CollectedGoFiles {
		cmd := exec.Command("goimports")
		cmd.Args = append(cmd.Args, args...)
		cmd.Args = append(cmd.Args, pth)
		cmd.Env = os.Environ()
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
