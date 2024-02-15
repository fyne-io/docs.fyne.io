---
layout: page
tags: [api]
title: Upgrading to v2.0
---

As our first ["major"](https://semver.org/) release the API updates in 2.0
contain some breaking changes.
Upgrading to v2.0.0 may break some applications so we recommend reading this
document carefully before performing the update.

To opt into these changes you will need to use the new import path of `fyne.io/fyne/v2`. This means updating all of your application imports, when you are ready.

## Modules

If your project has a `go.mod` file then you can edit the `require` line to use
version `v2.0.0`, or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne/v2@v2.0.0
```

If the above command did not automatically re-write your import declarations you
will need to change `fyne.io/fyne` to `fyne.io/fyne/v2` in all import lines.
The next time you build or run your app it will be using the 2.0 release.

## GOPATH

If you are not using modules then you will need to update the Fyne checkout in
your go source code. To do this execute the following command:

```bash
go get -u fyne.io/fyne/v2
```

You will need to change `fyne.io/fyne` to `fyne.io/fyne/v2` in all import lines.
Any apps without a module file will now use the 2.0 release.

## Changes

After updating all of the import paths to `fyne.io/fyne/v2` you may find that 
certain code does not work. This is more likely to be true if you are upgrading from
a version prior to the v1.4.0 release.

### Breaking changes

These items must be addressed when upgrading for your app to compile and run:

* Coordinate system to float32 - usage of `Size`, `Position` may need to be updated to `float32`
* `DraggedEvent` and `ScrollEvent` moved from (`int`, `int`) to `Delta` (using `float32`)

* There is a new theme API, to use your current theme code change from `SetTheme(myTheme)` to `SetTheme(theme.FromLegacy(myTheme))`
* Usage of `theme.NewThemedResource` should remove second, ignored, parameter
* Old button code using `Style` or `HideShadow` should now use `Importance`
* The `WidgetRenderer.BackgroundColor` method was removed
  - to add backgrounds you should now place rectangles under the widget

* The following deprecated types were removed, please move to replacements
  - `dialog.FileIcon` (now `widget.FileIcon`)
  - `widget.Radio` (now `widget.RadioGroup`)
  - `widget.AccordionContainer` (now `widget.Accordion`)
  - `layout.NewFixedGridLayout()` (now `layout.NewGridWrapLayout()`)
  - `widget.ScrollContainer` (now `container.Scroll`)
  - `widget.SplitContainer` (now `container.Spilt`)
  - `widget.Group` (replaced by `widget.Card`)
  - `widget.Box` (now `container.NewH/VBox` with `Children` field moved to `Objects`)
  - `widget.TabContainer` and `widget.AppTabs` (now `container.AppTabs`)

* A commonly used `widget.Refresh` call was finally removed, use `Widget.Refresh()` instead

### Other changes

These changes could cause behaviour inconsistency, so be sure to test your app
after performing the upgrade

* Dialogs no longer show when created, unless using the ShowXxx convenience methods
* Entry widget now contains scrolling so should no longer be wrapped in a scroll container
* iOS apps preferences will be lost in this upgrade as we move to more advanced storage


