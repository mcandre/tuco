package mageextras

import (
	"github.com/magefile/mage/mg"
)

// GoLint runs golint.
func GoLint(args ...string) error {
	mg.Deps(CollectGoFiles)

	for pth := range CollectedGoFiles {
		var as []string
		as = append(as, args...)
		as = append(as, pth)

		if err := Run("golint", as...); err != nil {
			return err
		}
	}

	return nil
}
