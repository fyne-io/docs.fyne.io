---
title: Rectangle

weight: 3010

aliases:
- /canvas/rectangle.html
- /tour/canvas/rectangle
- /tour/canvas/
- /canvas/
---

`canvas.Rectangle` is the simplest canvas object in Fyne. It displays
a block of the specified colour. You can also set the colour using
the `FillColor` field.

In this example the rectangle fills the window as it is
the only content element.

```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Rectangle")

	rect := canvas.NewRectangle(color.White)
	w.SetContent(rect)

	w.Resize(fyne.NewSize(150, 100))
	w.ShowAndRun()
}
```

Other `fyne.CanvasObject` types have more configuration, let us look
[next](text) at `canvas.Text`.
