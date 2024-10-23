---
title: Image

redirect_from:
  - /tour/canvas/image
---

A `canvas.Image` represents a scalable image resource in Fyne.
It can be loaded from a resource (as shown in the example), from an
image file, from a URI location containing an image, from an `io.Reader`,
or from a Go `image.Image` in memory.

Images can be bitmap based (like PNG and JPEG) or vector based
(such as SVG). Where possible we recommend scalable images as they will
continue to render well as sizes change.

## Image from a resource

The most common approach is to create a `canvas.Image` from a `fyne.Resource`.

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

	r := theme.AccountIcon()
	image := canvas.NewImageFromResource(r)
	image.SetMinSize(fyne.Size{Width: 200, Height: 200})
	w.SetContent(image)

	w.ShowAndRun()
}
```

In our example we use `theme.AccountIcon()`, because it is already available in the Fyne library. To use your own images you can bundle them into your project, which will make them available as Fyne resources. For more information please see [Bundling resources]({% link extend/bundle.md %}).

## Image from a URL

Images can also be created from URLs with `canvas.NewImageFromURI()`. This will fetch the image during runtime from a public image URL.

Compared to the first example, an additional step is required to parse the URL into an URI:

```go
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")

	u, _ := storage.ParseURI("https://avatars.githubusercontent.com/u/36045855")
	image := canvas.NewImageFromURI(u)
	image.SetMinSize(fyne.Size{Width: 200, Height: 200})
	w.SetContent(image)

	w.ShowAndRun()
}
```

Please note that is it often better to bundle images with your Fyne app. Your image will be rendered faster and you avoid potential issues related to network access during runtime.

## Image from other sources

You can also create an image from other sources:

- `image.Image`: `canvas.NewImageFromImage()`
- `io.Reader`: `canvas.NewImageFromReader()`
- file: `canvas.NewImageFromFile()`

For more information please see the API documentation for `canvas.Image`.

## Fillmode

You can specify how the images should expand to fill or fit the available space with `Image.Fillmode`.

The default fill mode is `canvas.ImageFillStretch` which will cause it
to fill the space specified (through `Resize()` or layout).

Alternatively you could use `canvas.ImageFillContain` to ensure that
the aspect ratio is maintained and the image is within the bounds.

When using `canvas.ImageFillStretch` or `canvas.ImageFillContain` you need to
specify a minimum size for your image with `Image.SetMinSize()`.
This ensures that your image is always rendered with the same size on different platforms and screen sizes.

Another available fill mode is `canvas.ImageFillOriginal`.
This mode ensures that the image will have a minimum size equal to that of the original image size
and therefore does not require you to set a minimum size.
Be careful when using original image sizes as they may not
behave exactly as expected with different user interface scales.
As Fyne allows the entire user interface to scale, a 25px image file
may not be the same height as a 25 height fyne object.
