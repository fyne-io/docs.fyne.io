---
title: Grid Wrap
slug: gridwrap

weight: 4030

aliases:
- /container/gridwrap.html
- /tour/layout/gridwraplayout
---

Like the previous grid layout, the grid wrap layout creates an arrangement
of elements in a grid pattern. However this grid does not have a set
number of columns, instead it uses a fixed size for each cell and
then flows the content to as many rows as is needed to display the items.

You create a grid wrap layout using `layout.NewGridWrapLayout(size)`
where size specifies the size to apply to all child elements.
This layout is then passed as the first parameter to
`container.New(...)`.
The number of columns and rows will be calculated based on the current
size of the container.

Initially a grid wrap layout will have a single column, if you resize it
(as illustrated in the code comment to the right) it will rearrange
the child elements to fill the space.

```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
		text1, text2, text3)
	myWindow.SetContent(grid)

	// myWindow.Resize(fyne.NewSize(180, 75))
	myWindow.ShowAndRun()
}
```
