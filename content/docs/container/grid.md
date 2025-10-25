---
title: Grid

weight: 4020

aliases:
- /container/grid.html
- /tour/layout/gridlayout
---

The grid layout lays out the elements of a container in a grid pattern
with a fixed number of columns. Items will fill a single row until the
number of columns is met, after this a new row will be created.
Vertical space will be split equally between each of the rows of objects.

You create a grid layout using `layout.NewGridLayout(cols)` where cols
is the number of items (columns) you wish to have in each row. This
layout is then passed as the first parameter to
`container.New(...)`.

If you resize the container then each of the cells will resize equally
to share the available space.

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
	myWindow := myApp.NewWindow("Grid Layout")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}
```
