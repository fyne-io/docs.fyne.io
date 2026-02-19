---
title: Stack

weight: 4070

aliases:
- /container/stack.html
- /tour/layout/maxlayout
- /container/max
---

The [layout.NewStackLayout()](/api/v2/layout/pkg/#func--newstacklayout) is the simplest layout, it sets all items in
the container to be the same size as the container. This is not
often useful in general containers but can be suitable when composing
widgets.

The stack layout will require the container to be at least the size of the
largest item's minimum size. The objects will be drawn in the order
they are passed to the container, with the last being drawn top-most.

```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Stack Layout")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewStackLayout(), img, text)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
```
