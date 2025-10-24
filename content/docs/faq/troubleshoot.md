---
title: Troubleshooting
slug: troubleshoot

order: 10030
---

Some things can not go as expected during setup or in compiling your first app. We try to address these problems here.
Remember you can also check your configuration using the
[Fyne Setup](https://geoffrey-artefacts.fynelabs.com/github/andydotxyz/fyne-io/setup/latest/) tool.

## Compiler Issues

**Q: command not found: fyne**

**A:** If you have installed the Fyne command using `go install fyne.io/tools/cmd/fyne@latest` but you
see an error when trying to run it then the most likely issue is that your Go install has not
set up your PATH environment variable correctly.

Go will install tools to the `bin` directory inside the user's `GOPATH` location
(normally `~/go`) - you can check this by seeing if your `PATH` variable includes this location.
If it seems to be missing then you should update your `PATH` environment variable to include
`~/go/bin`, or if you have changed install location then you can use `$(go env GOPATH)/bin` instead.

**Q: build constraints exclude all Go files in ...**

**A:** If you are cross compiling you may see an error about go files being excluded, followed
by a build failure. When doing a standard Go cross-compile it will automatically turn off CGo.
To fix this be sure to set `CGO_ENABLED=1` in your compile command.


**Q: cc1.exe: sorry, unimplemented: 64-bit mode not compiled in**

**A:** Windows compilation will sometimes complain that 64 bit mode is not available.
This is normally because the wrong compiler is installed, or the configuration is incorrect.
If you have followed our [installation instructions](/started/) for MSYS2 and MingW64 then make sure
you are using the start menu launcher titled "MSYS2 MinGW 64-bit".

## Distribution

**Q: Apple macOS says my app is damaged when it is downloaded**

**A:** When files are downloaded on a macOS computer they are marked with a "quarantine" flag so they are checked by the OS for problems. If your application is signed with a certificate purchased from Apple this is not a problem.
However if you want to share your software without that cost this error may appear - and on M1/2 computers it is not possible to use the *System Settings* to allow the app to run.

The fix is to remove the quarantine flag, which you can do by opening the *Terminal* and executing the following command:

    sudo xattr -r -d com.apple.quarantine MyApp.app
