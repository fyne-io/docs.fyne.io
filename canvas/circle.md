---
title: Circle

redirect_from:
  - /tour/canvas/circle
---

`canvas.Circle` defines a circle shape filled by the specified
colour. You may also set a `StrokeWidth` and therefore a different
`StrokeColor` as shown in this example.

The circle will fill the space specified by calling `Resize()` or
by the layout it is controlled by. As the example sets the circle
as the window content it will resize to fill the window, within a
basic padding (controlled by the theme).

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
	w := myApp.NewWindow("Circle")

	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = color.Gray{Y: 0x99}
	circle.StrokeWidth = 5
	w.SetContent(circle)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
```

All these have been basic types that can be rendered by our driver
with no additional information. Next we will look at more complex types
starting with [`Image`](image).
