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
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")

	r := theme.AccountIcon()
	image := canvas.NewImageFromResource(r)
	w.SetContent(image)
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
```

In our example we use `theme.AccountIcon()`, because it is already available in the Fyne library. To use your own images you can bundle them into your project, which will make them available as Fyne resources. For more information please see [Bundling resources]({% link extend/bundle.md %}).

Please note that we are setting a size for the window. This is necessary, because images normally do not have a minimum size and Fyne would draw a window with zero width and height. For more details please see the section about [Fillmode](#fillmode).

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
	w.SetContent(image)
	w.Resize(fyne.NewSize(200, 200))
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

Please note that with `canvas.ImageFillStretch` and `canvas.ImageFillContain` images
will use the available space of their container, but do not have a minimum size.
This means that layouts that shrink their objects to their minimum size
(e.g. `layout.NewHBoxLayout()`) will shrink images to zero.
In those cases you can set a minimum size yourself with `Image.SetMinSize()`.

Another available fill mode is `canvas.ImageFillOriginal`.
It ensures that the container grows to the pixel dimensions required to fit the original image.
The aspect of the image will be maintained so,
as with ImageFillContain there may be transparent areas around the image.

Be careful when using original image sizes as they may not
behave exactly as expected with different user interface scales.
As Fyne allows the entire user interface to scale, a 25px image file
may not be the same height as a 25 height fyne object.
