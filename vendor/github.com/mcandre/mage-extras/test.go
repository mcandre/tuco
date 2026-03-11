package mageextras

// UnitTest executes the Go unit test suite.
func UnitTest(args ...string) error {
	var as []string
	as = append(as, "test")
	as = append(as, args...)
	return Run("go", as...)
}
