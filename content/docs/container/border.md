---
title: Border

weight: 4040

aliases:
- /tour/layout/borderlayout
---

The border layout is possibly the most widely used to construct user
interfaces as it allows the positioning of items around a central element
which will expand to fill the space. To create a border container you need
to pass the `fyne.CanvasObject`s that should be positioned in a border
position to the constructor's first four parameters. This
syntax is basically just
`container.NewBorder(top, bottom, left, right, center)` as illustrated in
the example.

Any items passed to the container after the first four items will be
positioned to the central area and will expand to fill the space available.
You can also pass `nil` to border parameters that you wish to leave empty.

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
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("content", color.White)
	content := container.NewBorder(top, nil, left, nil, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
```

Note that all items in the center will expand to fill the space (as if
they were in a [`layout.MaxLayout`](/container/max) container).
To manage the area yourself you can use any `fyne.Container` as the
content instead.
