# tuco: minimalist Go port multiplexer

[![API Docs](https://badgen.net/static/docs/API/007d9c?icon=go)](https://pkg.go.dev/github.com/mcandre/tuco) [![CI/CD](https://badgen.net/github/checks/mcandre/tuco?icon=github)](https://github.com/mcandre/tuco/actions) [![License](https://badgen.net/static/license/BSD/ab2b29?icon=freebsd)](LICENSE.md)

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

# DOWNLOAD

```sh
go install github.com/mcandre/tuco/cmd/tuco@latest
```

## Prerequisites

* [Go](https://go.dev/)

For details on tuning tuco, see [CONFIGURATION](CONFIGURATION.md).

For details on building from source, see [DEVELOPMENT](DEVELOPMENT.md).

# ABOUT

tuco automates more low level steps involved in managing crosscompilation for Go projects. So that you can focus on developing your application.

# FEATURES

* Parallelism
* IaC friendly
* Easy port selection with YAML comment toggles
* Few dependencies
* Automatically corrects chmod bits inside tarballs
* Logical directory structure for straightforward binary based OS packaging

# RESOURCES

Prior art, personal plugs, and tools for developing portable applications (including non-Go projects)!

* [goreleaser](https://goreleaser.com/) - opinionated Go multiplexer
* [mcandre/crit](https://github.com/mcandre/crit) - Rust multiplexer
* [mcandre/rockhopper](https://github.com/mcandre/rockhopper) - OS package multiplexer
* [tree](https://en.wikipedia.org/wiki/Tree_(command)) - an CLI file manager
* [xgo](https://github.com/techknowlogick/xgo) - cGo multiplexer

🐹🍹
