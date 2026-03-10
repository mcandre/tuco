# INSTALL GUIDE

In addition to Go modules, tuco also supports alternative installation methods.

# INSTALL (PRECOMPILED BINARIES)

Precompiled binaries may be installed manually.

## Install

1. Download a [tarball](https://github.com/mcandre/tuco/releases) corresponding to your environment's architecture and OS.
2. Extract executables into a selected directory.

   Examples:

   * `~/.local/bin` (XDG compliant per-user)
   * `/usr/local/bin` (XDG compliant global)
   * `~/bin` (BSD)
   * `~\AppData\Local` (native Windows)

## Postinstall

Ensure the selected directory is registered with your shell's `PATH` environment variable.

## Uninstall

Remove the application executables from the selected directory.

## System Requirements

### Bitness

64

### Hosts

* AIX (PowerPC)
* DragonflyBSD (Intel)
* FreeBSD (ARM, Intel)
* Illumos (Intel)
* Linux (ARM, Intel, LoongArch, MIPS, MIPS LE, PowerPC, PowerPC LE, RISC-V, s390x)
* macOS 26 Tahoe+ (ARM, Intel)
* NetBSD (ARM, Intel)
* OpenBSD (ARM, Intel, PowerPC, RISC-V)
* Plan9 (Intel)
* Solaris (Intel)
* Windows 11+ (ARM, Intel)

### Prerequisites

* [Go](https://go.dev/)

# INSTALL (COMPILE FROM SOURCE)

```sh
git clone https://github.com/mcandre/tuco.git
cd tuco
go install ./...
```

## System Requirements

### Prerequisites

* [git](https://git-scm.com/)
* [Go](https://go.dev/)

For more information on developing tuco, see our [development guide](DEVELOPMENT.md).
