---
title: Image

weight: 3050

aliases:
- /canvas/image.html
- /tour/canvas/image
---

A [canvas.Image](/api/v2/canvas/image/#type--image) represents a scalable image resource in Fyne.
It can be loaded from a resource (as shown in the example), from an
image file, from a URI location containing an image, from an `io.Reader`, or from a Go `image.Image` in memory.

The default image fill mode is `canvas.ImageFillStretch` which will
cause it to fill the space specified (through `Resize()` or layout).
Alternatively you could use `canvas.ImageFillContain` to ensure that
the aspect ratio is maintained and the image is within the bounds.
Further to this you can use `canvas.ImageFillOriginal` (as used
in the example here) which ensures that it will also have a minimum size
equal to that of the original image size.

```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")

	image := canvas.NewImageFromResource(theme.FyneLogo())
	// image := canvas.NewImageFromURI(uri)
	// image := canvas.NewImageFromImage(src)
	// image := canvas.NewImageFromReader(reader, name)
	// image := canvas.NewImageFromFile(fileName)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.ShowAndRun()
}
```

Images can be bitmap based (like PNG and JPEG) or vector based
(such as SVG). Where possible we recommend scalable images as they will
continue to render well as sizes change.
Be careful when using original image sizes as they may not
behave exactly as expected with different user interface scales.
As Fyne allows the entire user interface to scale, a 25px image file
may not be the same height as a 25 height fyne object.
