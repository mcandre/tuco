# tuco: minimalist Go port multiplexer

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/mcandre/tuco) [![Test](https://github.com/mcandre/tuco/actions/workflows/test.yml/badge.svg)](https://github.com/mcandre/tuco/actions/workflows/test.yml) [![license](https://img.shields.io/badge/license-BSD-0)](LICENSE.md) [![Donate](https://img.shields.io/badge/%E2%99%A5-Sponsor-BF3988)](https://github.com/sponsors/mcandre)

![la tuza](tuco.png)

# SUMMARY

tuco streamlines Go application porting.

# EXAMPLE

```console
% cd example

% tuco

% ls bin/hello-ports
hello-aix-ppc64.tgz
hello-darwin-amd64.tgz
hello-darwin-arm64.tgz
...
```

For more CLI option, run `tuco -help`.

For more ports, run `go tool dist list`.

# DOWNLOAD

```sh
go install github.com/mcandre/tuco/cmd/tuco@latest
```

## Prerequisites

* [Go](https://go.dev/)

For more information on developing tuco, see our [development guide](DEVELOPMENT.md).

# ABOUT

tuco automates more low level steps involved in managing crosscompilation for Go projects. So that you can focus on developing your application.

# FEATURES

* Parallelism
* IaC friendly
* Easy port selection with YAML comment toggles
* Few dependencies
* Automatically corrects chmod bits inside tarballs
* Logical directory structure for straightforward binary based OS packaging

# CONFIGURATION

For details on tuning tuco, see our [configuration guide](CONFIGURATION.md).

# RESOURCES

Prior art, personal plugs, and tools for developing portable applications (including non-Go projects)!

* [goreleaser](https://goreleaser.com/) - opinionated Go multiplexer
* [mcandre/crit](https://github.com/mcandre/crit) - Rust multiplexer
* [mcandre/rockhopper](https://github.com/mcandre/rockhopper) - OS package multiplexer
* [tree](https://en.wikipedia.org/wiki/Tree_(command)) - an CLI file manager
* [xgo](https://github.com/techknowlogick/xgo) - cGo multiplexer

🐹🍹
