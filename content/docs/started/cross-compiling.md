---
layout: page
title: Compiling for different platforms
slug: cross-compiling

weight: 1140
---

Cross compiling with Go is designed to be simple - we just set the environment variable `GOOS` for the target Operating System (and `GOARCH` if targeting a different architecture). Unfortunately when using native graphics calls the use of CGo in Fyne makes this a little harder.

### Compiling from a development computer

To cross-compile a Fyne application you will also have to set `CGO_ENABLED=1` which tells go to enable the C compiler (this is normally turned off when the target platform is different to the current system). Doing so unfortunately means that you must have a C compiler for the target platform that you are going to compile for.
After installing the appropriate compilers you will also need to set the `CC` environment variable to tell Go which compiler to use.

There are many ways to install the required tools - and different tools that can be used. The configuration recommended by the Fyne developers is:

| GOOS (target) | CC | provider | download | notes |
|------|----|----------|----------|-------|
| `darwin`  | `o32-clang` | osxcross | [from github.com](https://github.com/tpoechtrager/osxcross) | You will also need to install the macOS SDK (instructions at the download link) |
| `windows` | `x86_64-w64-mingw64-gcc` | mingw64 | package manager | For macOS use [homebrew](https://brew.sh) |
| `linux`   | `gcc` or `x86_64-linux-musl-gcc` | gcc or musl-cross | [cygwin](https://www.cygwin.com/) or package manager | musl-cross is available from [homebrew](https://brew.sh) to provide the linux gcc. You will also need to install X11 and mesa headers for compilation. |

With the environment variables above set you should be able to compile in the usual manner.
If further errors occur it is likely to be due to missing packages. Some target platforms require additional libraries or headers to be installed for the compilation to succeed.

### Using a virtual environment

As a Linux system is able to cross compile to macOS and Windows easily it can be simpler to use a virtualised environment when you are not developing from Linux. Docker images are a useful tool for a complex build configuration and this works for Fyne as well. There are different tools that can be used. The tool recommended by the Fyne developers is [fyne-cross](https://github.com/fyne-io/fyne-cross). It has been inspired by [xgo](https://github.com/karalabe/xgo) and uses a [docker image](https://hub.docker.com/r/fyneio/fyne-cross) built on top of the [golang-cross](https://github.com/docker/golang-cross) image,
that includes the MinGW compiler for windows, and a macOS SDK, along with the Fyne requirements.

fyne-cross allows to build binaries and create distribution packages for the following targets:

| GOOS | GOARCH |
|------|----|
| darwin | amd64 |
| darwin | 386 |
| linux | amd64 |
| linux | 386 |
| linux | arm64 |
| linux | arm |
| windows | amd64 |
| windows | 386 |
| android | amd64 |
| android | 386 |
| android | arm64 |
| android | arm |
| ios | |
| freebsd | amd64 |
| freebsd | arm64 |

> Note: iOS compilation is supported only on darwin hosts.

#### Requirements

- go >= 1.13
- docker

#### Installation

You can install `fyne-cross` using the following command (requires Go 1.16 or later):

```
go install github.com/fyne-io/fyne-cross@latest
```

For earlier versions of Go, you need to use the following command instead:

```
go get github.com/fyne-io/fyne-cross
```

#### Usage

```
fyne-cross <command> [options]

The commands are:

	darwin        Build and package a fyne application for the darwin OS
	linux         Build and package a fyne application for the linux OS
	windows       Build and package a fyne application for the windows OS
	android       Build and package a fyne application for the android OS
	ios           Build and package a fyne application for the iOS OS
	freebsd       Build and package a fyne application for the freebsd OS
	version       Print the fyne-cross version information

Use "fyne-cross <command> -help" for more information about a command.
```

#### Wildcards

The `arch` flag support wildcards in case want to compile against all supported GOARCH for a specified GOOS

Example:

```
fyne-cross windows -arch=*
```

is equivalent to

```
fyne-cross windows -arch=amd64,386
```

#### Example

The example below cross compile and package the [fyne examples application](https://github.com/fyne-io/examples)

```
git clone https://github.com/fyne-io/examples.git
cd examples
```

##### Compile and package the main example app

```
fyne-cross linux
```

> Note: by default fyne-cross will compile the package into the current dir.
>
> The command above is equivalent to: `fyne-cross linux .`

##### Compile and package a particular example app

```
fyne-cross linux -output bugs ./cmd/bugs
```
