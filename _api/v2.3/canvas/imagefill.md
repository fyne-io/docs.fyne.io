---
layout: page
tags: [api]
title: Fyne API "canvas.ImageFill"
---

# canvas.ImageFill
---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type ImageFill

```go
type ImageFill int
```

ImageFill defines the different type of ways an image can stretch to fill its space.

```go
const (
	// ImageFillStretch will scale the image to match the Size() values.
	// This is the default and does not maintain aspect ratio.
	ImageFillStretch ImageFill = iota
	// ImageFillContain makes the image fit within the object Size(),
	// centrally and maintaining aspect ratio.
	// There may be transparent sections top and bottom or left and right.
	ImageFillContain // (Fit)
	// ImageFillOriginal ensures that the container grows to the pixel dimensions
	// required to fit the original image. The aspect of the image will be maintained so,
	// as with ImageFillContain there may be transparent areas around the image.
	// Note that the minSize may be smaller than the image dimensions if scale > 1.
	ImageFillOriginal
)
```
