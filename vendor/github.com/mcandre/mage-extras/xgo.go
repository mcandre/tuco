package mageextras

import (
	"os"
)

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo(outputPath string, args ...string) error {
	if err := os.MkdirAll(outputPath, os.ModeDir|0775); err != nil {
		return err
	}

	var as []string
	as = append(as, "-dest")
	as = append(as, outputPath)
	as = append(as, args...)
	return Run("xgo", as...)
}
