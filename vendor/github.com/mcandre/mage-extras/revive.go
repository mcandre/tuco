package mageextras

// Revive runs revive.
func Revive(args ...string) error {
	var as []string
	as = append(as, "-exclude", "vendor/...")
	as = append(as, args...)
	as = append(as, AllPackagesPath)
	return Run("revive", as...)
}
