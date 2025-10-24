---
title: Container and Layouts
slug: container

weight: 2020

aliases:
- /tour/basics/container
---

In the previous example we saw how to set a `CanvasObject` to the
content of a `Canvas`, but it is not very useful to only show
one visual element. To show more than one item we use the `Container` type.

As the `fyne.Container` also is a `fyne.CanvasObject`, we can set it to be
the content of a `fyne.Canvas`. In this example we create 3 text objects
and then place them in a container using the `container.NewWithoutLayout()` function.
As there is no layout set we can move the elements around like you see
with `text2.Move()`.

```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))
	content := container.NewWithoutLayout(text1, text2)
	// content := container.New(layout.NewGridLayout(2), text1, text2)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
```

A `fyne.Layout` implements a method for organising items within a container.
By uncommenting the `container.New()` line in this example you
alter the container to use a grid layout with 2 columns. Run this code
and try resizing the window to see how the layout automatically configures
the contents of the window. Notice also that the manual position
of `text2` is ignored by the layout code.

To see more you can check out the [`Layout list`](layouts).
