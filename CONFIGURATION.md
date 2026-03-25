# CONFIGURATION

tuco uses [YAML](https://yaml.org/) syntax for configuration files.

# tuco.yaml

tuco reads a configuration file `tuco.yaml` in the current working directory.

# `debug`

Default: `false`

When `true`, enables additional logs.

Example:

```yaml
debug: true
```

# `artifacts`

Default: `"bin"`

Customize the toplevel directory for binary artifacts.

Example:

```yaml
artifacts: "dist"
```

# `banner`

Required, nonblank.

Software application name.

Example:

```yaml
banner = "hello"
```

# `jobs`

Default: 4

Customize concurrent Go worker limit.

Example:

```yaml
jobs: 8
```

# `go_args`

Default: (empty)

Supply additional CLI arguments to `go build`... commands.

Example:

```yaml
go_args:
- "-v"
```

# Go Ports

Target Go [ports](https://go.dev/wiki/PortingPolicy).

To enumerate available ports, run `go tool dist list`.

## `os`

Enable GOOS values.

Example:

```yaml
# Skip mobile SDKs
os:
- "aix"
# - "android"
- "darwin"
- "dragonfly"
- "freebsd"
- "illumos"
# - "ios"
- "js"
- "linux"
- "netbsd"
- "openbsd"
- "plan9"
- "solaris"
- "wasip1"
- "windows"
```

## `arch`

Enable GOARCH values.

Example:

```yaml
arch:
- "386"
- "amd64"
- "arm"
- "arm64"
- "loong64"
- "mips"
- "mips64"
- "mips64le"
- "mipsle"
- "ppc64"
- "ppc64le"
- "riscv64"
- "s390x"
- "wasm"
```

# `port_excludes`

Default: (empty)

Skip named Go ports.

Useful for disabling broken or deprecated ports.

Example:

```yaml
port_excludes:
- "darwin/amd64"
```

# `archive_excludes`

Default:

```yaml
- ".DS_Store" # Finder
- "Thumbs.db" # Explorer
```

Skips corresponding file path patterns when archiving binaries.

Syntax: [Glob](https://pkg.go.dev/path/filepath#Match)

Example:

```yaml
excludes:
- ".DS_Store"  # Finder
- ".directory" # Dolpin
- "Thumbs.db"  # Explorer
```
