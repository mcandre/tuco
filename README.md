# tuco: Go port multiplexer

[![GitHub Downloads](https://img.shields.io/github/downloads/mcandre/tuco/total?logo=github)](https://github.com/mcandre/tuco/releases) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/mcandre/tuco) [![Test](https://github.com/mcandre/tuco/actions/workflows/test.yml/badge.svg)](https://github.com/mcandre/tuco/actions/workflows/test.yml) [![license](https://img.shields.io/badge/license-BSD-0)](LICENSE.md) [![Donate](https://img.shields.io/badge/-any?logo=gumroad&label=Donate&color=grey)](https://mcandre.gumroad.com/)

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

For more platforms and installation methods, see our [install guide](INSTALL.md).

# ABOUT

tuco elevates the Go compilation experience, taking care of many low level tasks. So tht you can focus on developing your appliaction.

# FEATURES

* Parallelism
* IaC friendly
* Easy port selection with YAML comment toggles
* Automatically corrects chmod bits inside tarballs
* Logical directory structure for straightforward binary based OS packaging

# CONFIGURATION

For details on tuning tuco, see our [configuration guide](CONFIGURATION).

# RESOURCES

Prior art, personal plugs, and tools for developing portable applications (including non-Go projects)!

* [mcandre/crit](https://github.com/mcandre/crit) - Rust multiplexer
* [mcandre/rockhopper](https://github.com/mcandre/rockhopper) - OS package multiplexer
* [tree](https://en.wikipedia.org/wiki/Tree_(command)) - an CLI file manager
* [xgo](https://github.com/techknowlogick/xgo) - cGo multiplexer

🐹🍹
