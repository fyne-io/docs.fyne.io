---
title: Raster

redirect_from:
  - /tour/canvas/raster
---

The `canvas.Raster` is like an image but draws exactly one spot
for each pixel on the screen. This means that as a user interface
scales or the image resizes more pixels will be requested to fill
the space. To do this we use a `Generator` function as illustrated in
this example - it will be used to return the colour of each pixel.

The generator functions can be pixel based (as in this example where we 
generate a new random colour for each pixel) or based on full images.
Generating complete images (with `canvas.NewRaster()`) is more efficient
but sometimes controlling pixels directly is more convenient.

```go
package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")

	raster := canvas.NewRasterWithPixels(
		func(_, _, w, h int) color.Color {
			return color.RGBA{R: uint8(rand.Intn(255)),
				G: uint8(rand.Intn(255)),
				B: uint8(rand.Intn(255)), A: 0xff}
		})
	// raster := canvas.NewRasterFromImage()
	w.SetContent(raster)
	w.Resize(fyne.NewSize(120, 100))
	w.ShowAndRun()
}
```

If your pixel data is stored in an image you can load it through the
`NewRasterFromImage()` function which will load the image to display
pixel perfect on screen.
