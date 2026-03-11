package mageextras

// Nakedret runs nakedret.
func Nakedret(args ...string) error {
	var as []string
	as = append(as, args...)
	as = append(as, AllPackagesPath)
	return Run("nakedret", as...)
}
