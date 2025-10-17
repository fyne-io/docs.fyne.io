---
layout: page
tags: [api]
title: Upgrading to v2.7
---

The 2.7 release is fully backward compatible with 2.6.3 and earlier.
There are no known changes to existing behaviour beyond bug fixes, so the upgrade should be trivial.

## Updating

Open your `go.mod` file and alter the the `require` line to use version `v2.7.0`,
or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne/v2@v2.7.0
```

The next time you build or run your app it will be using the 2.7 API,
showing the new curved corners in input widgets and selection.

## Fyne command

You should update the `fyne` tool to get some of the new features and bug fixes.
This is done by installing it from the tools repository.
You can make the upgrade by using a `go install` command similar to above:

```bash
go install fyne.io/tools/cmd/fyne@latest
```

After that completes, check you have the new version installed by running `fyne version`.

## Changes

This release is completely backwards compatible so your code will compile and
run as expected - but a lot faster due to recent performance improvements.

