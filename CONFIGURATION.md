# CONFIGURATION GUIDE

# FLAGS

## `-clean`

Remove artifacts.

## `-help`

Show usage menu.

## `-version`

Show version identifier.

# YAML

tuco looks for a configuration file `tuco.yaml` in the current working directory.

## `debug`

Default: `false`

When `true`, enables additional logs.

Example:

```yaml
debug: true
```

## `artifacts`

Default: `"bin"`

Customize the toplevel directory for binary artifacts.

Example:

```yaml
artifacts: "dist"
```

## `banner`

Required, nonblank.

Software application name.

Example:

```yaml
banner = "hello"
```

## `go_args`

Default: (empty)

Supply additional CLI arguments to `go build`... commands.

Example:

```yaml
go_args:
- "-v"
```

## `ports`

Target Go [ports](https://go.dev/wiki/PortingPolicy).

Example:

```yaml
ports:
- "darwin/amd64"
- "darwin/arm64"
- "linux/amd64"
- "linux/arm64"
- "linux/riscv64"
- "windows/amd64"
- "windows/arm64"
# ...
```

To enumerate available ports, run `go tool dist list`.
