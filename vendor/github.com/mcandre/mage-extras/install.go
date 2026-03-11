package mageextras

import (
	"os"
	"path"

	"github.com/magefile/mage/mg"
)

// Install builds and installs Go applications.
func Install(args ...string) error {
	var as []string
	as = append(as, "install")
	as = append(as, args...)
	as = append(as, AllPackagesPath)
	return Run("go", as...)
}

// Uninstall deletes installed Go applications.
func Uninstall(applications ...string) error {
	mg.Deps(LoadGoBinariesPath)

	for _, application := range applications {
		if err := os.RemoveAll(path.Join(LoadedGoBinariesPath, application)); err != nil {
			return err
		}
	}

	return nil
}
