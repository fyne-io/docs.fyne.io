---
layout: page
tags: [api]
title: Upgrading to v1.3
---

The 1.3 release is fully backward compatible with 1.2.4 and earlier, so upgrading
is as simple as updating the version of code you compile with.
This is different depending on whether or not you use go modules.

## Modules

If your project has a `go.mod` file then you can edit the `require` line to use
version `v1.3.0`, or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne@v1.3.0
```

The next time you build or run your app it will be using the 1.3 release

## GOPATH

If you are not using modules then you will need to update the Fyne checkout in
your go source code. To do this execute the following command:

```bash
go get -u fyne.io/fyne
```

Any apps without a module file will now use the 1.3 release.

## Changes

Although this release is backwards compatible so your code will compile and
run as expected, there are some changes which you may notice.

* Running `Run()` or `ShowAndRun()` in a goroutine is not supported and will now panic
* Running on Linux applications are now 20% smaller, better matching other OS
* Resizing windows now live-refreshes the content and animations will continue
* `FixedGridLayout` has been renamed `GridWrapLayout` - the old APIs remain but deprecated
* Some other APIs have been deprecated and will be removed in 2.0 later in 2020

