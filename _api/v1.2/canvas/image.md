---
layout: page
tags: [api]
title: Fyne API canvas
---

# canvas
---
```go
import "fyne.io/fyne/canvas"
```

## Usage

#### type Image

```go
type Image struct {

	// one of the following sources will provide our image data
	File     string        // Load the image from a file
	Resource fyne.Resource // Load the image from an in-memory resource
	Image    image.Image   // Specify a loaded image to use in this canvas object

	Translucency float64   // Set a translucency value > 0.0 to fade the image
	FillMode     ImageFill // Specify how the image should scale to fill or fit
}
```

Image describes a drawable image area that can render in a Fyne canvas The image may be a vector or a bitmap representation and it will fill the area. The fill mode can be changed by setting FillMode to a different ImageFill.

#### func  NewImageFromFile

```go
func NewImageFromFile(file string) *Image
```
NewImageFromFile creates a new image from a local file. Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.

#### func  NewImageFromImage

```go
func NewImageFromImage(img image.Image) *Image
```
NewImageFromImage returns a new Image instance that is rendered from the Go image.Image passed in. Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.

#### func  NewImageFromResource

```go
func NewImageFromResource(res fyne.Resource) *Image
```
NewImageFromResource creates a new image by loading the specified resource. Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.

#### func (*Image) Alpha

```go
func (i *Image) Alpha() float64
```
Alpha is a convenience function that returns the alpha value for an image based on its Translucency value. The result is 1.0 - Translucency.

#### func (*Image) Hide

```go
func (r *Image) Hide()
```
Hide will set this object to not be visible

#### func (*Image) MinSize

```go
func (r *Image) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise

#### func (*Image) Move

```go
func (r *Image) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*Image) Position

```go
func (r *Image) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to its parent / canvas

#### func (*Image) Refresh

```go
func (i *Image) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Image) Resize

```go
func (r *Image) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*Image) SetMinSize

```go
func (r *Image) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*Image) Show

```go
func (r *Image) Show()
```
Show will set this object to be visible

#### func (*Image) Size

```go
func (r *Image) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*Image) Visible

```go
func (r *Image) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise
