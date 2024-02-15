---
layout: page
tags: [api]
title: Upgrading to v1.4
---

The 1.4 release is fully backward compatible with 1.3.3 and earlier, so upgrading
is as simple as updating the version of code you compile with.
This is different depending on whether or not you use go modules.

## Modules

If your project has a `go.mod` file then you can edit the `require` line to use
version `v1.4.0`, or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne@v1.4.0
```

The next time you build or run your app it will be using the 1.4 release

## GOPATH

If you are not using modules then you will need to update the Fyne checkout in
your go source code. To do this execute the following command:

```bash
go get -u fyne.io/fyne
```

Any apps without a module file will now use the 1.4 release.

## Changes

Although this release is backwards compatible so your code will compile and
run as expected, there are some changes which you may notice.

* The main theme has changed slightly. The blue primary colour is different and users
can now choose from a range of primary colors.
* Buttons are now transparent with an outline by default, this may impact your app design.
* The Group widget has been deprecated and replaced with the new Card widget.
* Many container widgets have been deprecated and they have been replaced by new types in the `container` package.

