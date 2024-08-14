---
layout: page
tags: [api]
title: Upgrading to v2.4
---

The 2.4 release is fully backward compatible with 2.3.5 and earlier, so upgrading
is as simple as updating the version of code you compile with.
From thie v2.4.0 release of Fyne it requires Go 1.17, and so all projects are assumed to be using Go modules.

## Updating

Open your `go.mod` file and alter the the `require` line to use version `v2.4.0`,
or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne/v2@v2.4.0
```

The next time you build or run your app it will be using the 2.4 API,
showing the new curved corners in input widgets and selection.

## Fyne command

You should update the `fyne` tool for v2.4.0 to get some of the new features and bug fixes.
You can make the upgrade by using the `go get` command similarly to above:

```bash
go install fyne.io/fyne/v2/cmd/fyne@v2.4.0
```

After that completes, check you have the new version installed by running `fyne version`.

## Changes

Although this release is backwards compatible so your code will compile and
run as expected, there are some changes which you may notice.

* Fyne now requires Go 1.17 - be sure to upgrade if you were on an earlier version
* Using `fyne.TextWrapOff` as `Entry.Wrapping` to disable scrolling was confusing, instead you can now set `Entry.Scroll` to `container.ScrollNone`).
* Refreshing an image will now happen in the current goroutine not render process, apps may wish to add async image load
* Icons for macOS bundles are now padded and rounded, disable with "-use-raw-icon"
* Accordion widget now fills available space - put it inside a `VBox` container for old behavior (#4126)
