# DEVELOPMENT GUIDE

tuco follows standard, Go based operations for compiling and unit testing.

For advanced operations, such as linting, we further supplement with some software industry tools.

# BUILDTIME REQUIREMENTS

* a UNIX-like environment (e.g. [WSL](https://learn.microsoft.com/en-us/windows/wsl/))
* [Go](https://go.dev/)
* POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)
* Provision additional dev tools with `make`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.18 (run `asdf reshim` after provisioning)

# AUDIT

```sh
mage audit
```

# INSTALL

```sh
mage install
```

# UNINSTALL

```sh
mage uninstall
```

# LINT

```sh
mage lint
```

# TEST

```sh
mage test
```

# CROSSCOMPILE & ARCHIVE BINARIES

```sh
mage tuco
```

# CLEAN

```sh
mage clean
```
