---
title: Gradient

weight: 3070

since: 1.1

aliases:
- /canvas/gradient.html
- /tour/canvas/gradient
---

The last canvas primitive type is Gradient, available as
[canvas.LinearGradient](/api/v2/canvas/lineargradient/#type--lineargradient) and [canvas.RadialGradient](/api/v2/canvas/radialgradient/#type--radialgradient) which is used
to draw a gradient from one colour to another in various patterns.
You can create gradients using `NewHorizontalGradient()`,
`NewVerticalGradient()` or `NewRadialGradient()`.

To create a gradient you need a start and end colour - every colour
in between is calculated by the canvas. In this example we use 
`color.Transparent` to show how a gradient (or any other type) could
use an alpha value to be semi-transparent over the content behind.

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
	w := myApp.NewWindow("Gradient")

	gradient := canvas.NewHorizontalGradient(color.White, color.Transparent)
	//gradient := canvas.NewRadialGradient(color.White, color.Transparent)
	w.SetContent(gradient)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
```
