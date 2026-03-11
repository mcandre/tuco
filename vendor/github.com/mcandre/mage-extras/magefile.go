//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	mageextras "github.com/mcandre/mage-extras"
)

// Default references the default build task.
var Default = CoverageHTML

// CoverHTML denotes the HTML formatted coverage filename.
var CoverHTML = "cover.html"

// CoverProfile denotes the raw coverage data filename.
var CoverProfile = "cover.out"

// Audit runs a security audit.
func Audit() error { return Govulncheck() }

// Clean deletes build artifacts.
func Clean() error { mg.Deps(CleanCoverage); return nil }

// CleanCoverage deletes coverage data.
func CleanCoverage() error {
	if err := os.RemoveAll(CoverHTML); err != nil {
		return err
	}

	return os.RemoveAll(CoverProfile)
}

// CoverageHTML generates HTML formatted coverage data.
func CoverageHTML() error {
	mg.Deps(CoverageProfile)
	return mageextras.CoverageHTML(CoverHTML, CoverProfile)
}

// CoverageProfile generates raw coverage data.
func CoverageProfile() error { return mageextras.CoverageProfile(CoverProfile) }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Run("errcheck", "-blank") }

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Run("govulncheck", "-scan", "package", "./...") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoImports)
	mg.Deps(GoLint)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Revive)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	return nil
}

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// NoVendor lists non-vendored Go source files.
func NoVendor() error {
	mg.Deps(mageextras.CollectGoFiles)

	for pth, _ := range mageextras.CollectedGoFiles {
		fmt.Println(pth)
	}

	return nil
}

// Revive runs revive.
func Revive() error { return mageextras.Revive() }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Run("staticcheck", "./...") }

// Test executes the unit test suite.
func Test() error { return mageextras.UnitTest() }
