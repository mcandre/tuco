//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	mageextras "github.com/mcandre/mage-extras"

	"os"
	"os/exec"
)

// Default references the default build task.
var Default = Test

// Audit runs security checks.
func Audit() error { return Govulncheck() }

// Clean removes artifacts.
func Clean() error { mg.Deps(CleanExample); return CleanArtifacts() }

// CleanArtifacts removes artifacts.
func CleanArtifacts() error { return mageextras.Tuco("-clean") }

// CleanEample removes artifacts from example projects.
func CleanExample() error {
	cmd := exec.Command("tuco", "-clean")
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Dir = "example"
	return cmd.Run()
}

// Deadcode runs deadcode.
func Deadcode() error { return mageextras.Deadcode("./...") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Govulncheck("-scan", "package", "./...") }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Lint runs the linter suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoImports)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	return nil
}

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck("./...") }

// Tuco builds crossplatform binaries and tarballs.
func Tuco() error { return mageextras.Tuco() }

// Test runs a test suite.
func Test() error { return mageextras.UnitTest() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("tuco") }
