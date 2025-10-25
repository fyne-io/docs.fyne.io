---
title: Canvas and CanvasObject
slug: canvas

weight: 2010

aliases:
- /explore/canvas.html
- /tour/basics/canvas
- /tour/basics/widget
- /explore/
---

In Fyne, a `Canvas` is the area within which an application is drawn.
Each window has a canvas, which you can access with `Window.Canvas()`,
but usually, you will find functions on `Window` that avoid directly accessing
the canvas.

Everything that can be drawn in Fyne is a type of `CanvasObject`.
The example here opens a new window and then shows different types of
primitive graphical elements by setting the content of the window canvas.
There are many ways that each type of object can be customised, as
shown with the text and circle examples.

As well as changing the content shown using `Canvas.SetContent()`, it is
possible to change the properties of existing canvas objects. For example, you
can change the `FillColor` of a rectangle and then call `rect.Refresh()` to update
its appearance. When performing such updates from a goroutine, you should
use the `fyne.Do` function to queue the updates safely on the main goroutine,
ensuring thread safety as required since Fyne v2.6.0.

```go
package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	myCanvas.SetContent(rect)

	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		fyne.Do(func() {
			rect.FillColor = green
			rect.Refresh()
		})
	}()

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}
```

We can draw many different drawing elements in the same way, such as circles and text.

```go
func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}
```

## Widget

A `fyne.Widget` is a special type of canvas object that has interactive elements
associated with it. In widgets, the logic is separate from the way that
it looks (also called the `WidgetRenderer`).

Widgets are also types of `CanvasObject`, so we can set the
content of our window to a single widget. See how we create a new
`widget.Entry` and set it as the content of the window in this example.

```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Widget")

	myWindow.SetContent(widget.NewEntry())
	myWindow.ShowAndRun()
}
```
