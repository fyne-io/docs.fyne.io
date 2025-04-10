---
layout: page
tags: [api]
title: Upgrading to v2.6
---

The 2.6 release is fully backward compatible with 2.5.5 and earlier, 
however there are notable changes in how it works internally that may impact some apps.

Changes:

  * All callbacks and widget functions are called on the same event thread
  * It now matters whether your code runs in the background or not, see below
  * Menu shortcuts will be called ahead of shortcuts on the current focused widget
  * The fyne command-line tool has been moved to the fyne.io/tools project

## Updating

Open your `go.mod` file and alter the the `require` line to use version `v2.6.0`,
or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne/v2@v2.6.0
```

The next time you build or run your app it will be using the 2.6 API,
showing the new curved corners in input widgets and selection.

## Fyne command

You should update the `fyne` tool for v2.6.0 to get some of the new features and bug fixes.
This is now done by installing it from the tools repository.
You can make the upgrade by using the `go get` command similarly to above:

```bash
go install fyne.io/tools/cmd/fyne@latest
```

After that completes, check you have the new version installed by running `fyne version`.

## Changes

Although this release is backwards compatible so your code will compile and
run as expected, there are some changes which you may need to consider.

### Thread handling

**Fyne now uses a single thread / goroutine for handling all user events and data processing**

This means that all callbacks will occur in serial, meaning there is a small possibility of user
code having a deadlock that was not present before.
If this occurs in your app double check that you are not holding internal locks (they should not
be required for ensuring widget data safety any more).

**Background access to Fyne APIs must be performed through fyne.Do[AndWait] functions**

Because all Fyne callbacks now happen on a single thread (for speed and safety) apps using goroutines need to be updated.
If your app launches a goroutine (using the `go` keyword) any Fyne APIs called from within the
new context must be wrapped in a new accessor.

This new API has two varieties - `fyne.Do` and `fyne.DoAndWait`, for most situations the first will be best.
The "AndWait" version should be used if your app must wait for the user interface to process the code passed in and update.
If your change is a simple update and you can "fire and forget" simply call `fyne.Do`, the order of operations is guaranteed and changes will happen within the next graphical update.

### Smaller changes

 * Global shortcuts can now be registered by adding a menu item in the app's MainMenu. Items in the menu will be actioned instead of sending the shortcut to a widget if there is a match. For a small number of apps this may cause a behaviour change - if so the menu item could check the current focused widget and forward the action.

