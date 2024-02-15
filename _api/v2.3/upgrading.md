---
layout: page
tags: [api]
title: Upgrading to v2.3
---

The 2.3 release is fully backward compatible with 2.2.4 and earlier, so upgrading
is as simple as updating the version of code you compile with.
This is different depending on whether or not you use go modules.

## Modules

If your project has a `go.mod` file then you can edit the `require` line to use
version `v2.3.0`, or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne/v2@v2.3.0
```

The next time you build or run your app it will be using the 2.3 release,
showing the updated theme.

## GOPATH

If you are not using modules then you will need to update the Fyne checkout in
your go source code. To do this execute the following command:

```bash
go get -u fyne.io/fyne/v2
```

Any apps without a module file will now use the 2.3 release.

## Fyne command

You should update the `fyne` tool for v2.3.0 to get the web latest metadata
support and bug fixes.
You can make the upgrade by using the `go get` command similarly to above:

```bash
go get -u fyne.io/fyne/v2/cmd/fyne@v2.3.0
```

After that completes, check you have the new version installed by running `fyne version`.

## Changes

Although this release is backwards compatible so your code will compile and
run as expected, there are some changes which you may notice.

* Buttons are no longer transparent in the default theme, so items drawn under
them will not show through. As this was most commonly to show a color of button
consider setting `Button.Importance` to new `ErrorImportance` or `WarningImportance` instead.

