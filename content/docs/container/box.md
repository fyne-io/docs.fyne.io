---
title: Box

weight: 4010

aliases:
- /tour/layout/boxlayout
- /container/
---

As discussed in [Container and Layouts](/explore/container) elements
within a container can be arranged using a layout. This section explores
the builtin layouts and how to use them.

The most commonly used layout is `layout.BoxLayout` and it has two variants,
horizontal and vertical. A box layout arranges all elements in a single
row or column with optional spaces to assist alignment.

A horizontal box layout, created with `layout.NewHBoxLayout()` creates
an arrangement of items in a single row. Every item in the box will
have its width set to its `MinSize().Width` and the height will be
equal for all items, the largest of all the `MinSize().Height` values.
The layout can be used in a container or you can use the box widget
`widget.NewHBox()`.

A vertical box layout is similar but it arranges items in a column.
Each item will have its height set to minimum and all the widths will
be equal, set to the largest of the minimum widths.

To create an expanding space between elements (so that some are left
aligned and the others right aligned, for example) add a `layout.NewSpacer()`
as one of the items. A spacer will expand to fill all available space.
Adding a spacer at the beginning of a vertical box layout will cause
all items to be bottom aligned. You can add one to the beginning and
end of a horizontal arrangement to create a center alignment.

```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	myWindow.ShowAndRun()
}
```
