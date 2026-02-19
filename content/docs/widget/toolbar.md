---
title: Toolbar

weight: 5070

aliases:
- /widget/toolbar.html
- /tour/widget/toolbar
---

The toolbar widget creates a row of action buttons using
icons to represent each. The `widget.NewToolbar(...)`
constructor function takes a list of [widget.ToolbarItem](/api/v2/widget/toolbaritem/#type--toolbaritem)
parameters. The builtin types of toolbar items are action,
separator and spacer.

The most used item is an action that is created using the
`widget.NewToolbarItemAction(..)` function. An action takes
two parameters, first being the icon resource to draw and
the latter is the `func()` to call when tapped. This creates
a standard toolbar button.

You can use [widget.NewToolbarSeparator()](/api/v2/widget/toolbarseparator/#func--newtoolbarseparator) to create a
small divider between items in a toolbar (usually a thin
vertical line). Lastly you can use [widget.NewToolbarSpacer()](/api/v2/widget/toolbarspacer/#func--newtoolbarspacer)
to create a flexible space between elements. This is most
useful to right align the toolbar items that are listed
after the spacer.

A toolbar should always be at the top of the content area
so it's normal to add it to a [fyne.Container](/api/v2/fyne/container/#type--container) using the
`layout.BorderLayout` to align it above other content.

```go
package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Toolbar Widget")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
```
