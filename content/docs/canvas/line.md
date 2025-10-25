---
title: Line

weight: 3030

aliases:
- /canvas/line.html
- /tour/canvas/line
---

The `canvas.Line` object draws a line from the `Position1` (default
is top, left) to `Position2` (default is bottom, right).
You specify its colour and can vary the stroke width which otherwise
defaults to `1`.

A line position can be manipulated using the `Position1` or `Position2`
fields or by using the `Move()` and `Resize()` functions. For example 
a 0 width area will show a vertical line whereas 0 height would be
horizontal.

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
	w := myApp.NewWindow("Line")

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	w.SetContent(line)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
```

Lines are typically used in a custom layout or controlled manually.
Unlike text they have no natural (minimum) size but can be
used to great effect in complex layouts.
