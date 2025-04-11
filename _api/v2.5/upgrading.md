---
layout: page
tags: [api]
title: Upgrading to v2.5
---

The 2.5 release is fully backward compatible with 2.4.5 and earlier, so upgrading
is as simple as updating the version of code you compile with.
As of release v2.5.0, Fyne requires Go 1.19, and so all projects are assumed to be using Go modules.

## Updating

Open your `go.mod` file and alter the the `require` line to use version `v2.5.0`,
or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne/v2@v2.5.0
```

The next time you build or run your app it will be using the 2.5 API,
showing the new curved corners in input widgets and selection.

## Fyne command

You should update the `fyne` tool for v2.5.0 to get some of the new features and bug fixes.
You can make the upgrade by using the `go get` command similarly to above:

```bash
go install fyne.io/fyne/v2/cmd/fyne@v2.5.0
```

After that completes, check you have the new version installed by running `fyne version`.

## Changes

Although this release is backwards compatible so your code will compile and
run as expected, there are some changes which you may notice.

* Fyne now requires Go 1.19 - be sure to upgrade if you were on an earlier version
* Widgets are now able to have a non-standard theme if placed inside a `ThemeOverride` container, so theme code should use `theme.ForWidget()` instead of the static `theme` helper methods
* Mobile apps using a software keyboard will now be adjusted in size to fit the space - consider adding scrolling to taller containers
