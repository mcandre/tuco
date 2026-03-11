package mageextras

import (
	"fmt"
)

// CoverageHTML generates HTML formatted coverage data.
func CoverageHTML(htmlFilename string, profileFilename string) error {
	return Run(
		"go",
		"tool",
		"cover",
		fmt.Sprintf("-html=%s", profileFilename),
		"-o",
		htmlFilename,
	)
}

// CoverageProfile generates raw coverage data.
func CoverageProfile(profileFilename string) error {
	return Run(
		"go",
		"test",
		fmt.Sprintf("-coverprofile=%s", profileFilename),
	)
}
