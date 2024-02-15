---
layout: page
tags: [api]
title: Fyne API "canvas.Raster"
---

# canvas.Raster
---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type Raster

```go
type Raster struct {

	// Render the raster image from code
	Generator func(w, h int) image.Image

	// Set a translucency value > 0.0 to fade the raster
	Translucency float64
	// Specify the type of scaling interpolation applied to the raster if it is not full-size
	// Since: 1.4.1
	ScaleMode ImageScale
}
```

Raster describes a raster image area that can render in a Fyne canvas

#### func  NewRaster

```go
func NewRaster(generate func(w, h int) image.Image) *Raster
```
NewRaster returns a new Image instance that is rendered dynamically using the specified generate function. Images returned from this method should draw dynamically to fill the width and height parameters passed to pixelColor.

#### func  NewRasterFromImage

```go
func NewRasterFromImage(img image.Image) *Raster
```
NewRasterFromImage returns a new Raster instance that is rendered from the Go image.Image passed in. Rasters returned from this method will map pixel for pixel to the screen starting img.Bounds().Min pixels from the top left of the canvas object. Truncates rather than scales the image. If smaller than the target space, the image will be padded with zero-pixels to the target size.

#### func  NewRasterWithPixels

```go
func NewRasterWithPixels(pixelColor func(x, y, w, h int) color.Color) *Raster
```
NewRasterWithPixels returns a new Image instance that is rendered dynamically by iterating over the specified pixelColor function for each x, y pixel. Images returned from this method should draw dynamically to fill the width and height parameters passed to pixelColor.

#### func (*Raster) Alpha

```go
func (r *Raster) Alpha() float64
```
Alpha is a convenience function that returns the alpha value for a raster based on its Translucency value. The result is 1.0 - Translucency.

#### func (*Raster) Hide

```go
func (o *Raster) Hide()
```
Hide will set this object to not be visible.

#### func (*Raster) MinSize

```go
func (o *Raster) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise.

#### func (*Raster) Move

```go
func (o *Raster) Move(pos fyne.Position)
```
Move the object to a new position, relative to its parent.

#### func (*Raster) Position

```go
func (o *Raster) Position() fyne.Position
```
Position gets the current position of this canvas object, relative to its parent.

#### func (*Raster) Refresh

```go
func (r *Raster) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Raster) Resize

```go
func (r *Raster) Resize(s fyne.Size)
```
Resize on a raster image causes the new size to be set and then calls Refresh. This causes the underlying data to be recalculated and a new output to be drawn.

#### func (*Raster) SetMinSize

```go
func (o *Raster) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be.

#### func (*Raster) Show

```go
func (o *Raster) Show()
```
Show will set this object to be visible.

#### func (*Raster) Size

```go
func (o *Raster) Size() fyne.Size
```
Size returns the current size of this canvas object.

#### func (*Raster) Visible

```go
func (o *Raster) Visible() bool
```
Visible returns true if this object is visible, false otherwise.
