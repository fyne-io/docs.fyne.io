---
layout: page
tags: [api]
title: Fyne API "canvas.Image"
---

# canvas.Image
---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type Image

```go
type Image struct {

	// one of the following sources will provide our image data
	File     string        // Load the image from a file
	Resource fyne.Resource // Load the image from an in-memory resource
	Image    image.Image   // Specify a loaded image to use in this canvas object

	Translucency float64    // Set a translucency value > 0.0 to fade the image
	FillMode     ImageFill  // Specify how the image should expand to fill or fit the available space
	ScaleMode    ImageScale // Specify the type of scaling interpolation applied to the image
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

#### func  NewImageFromReader

```go
func NewImageFromReader(read io.Reader, name string) *Image
```
NewImageFromReader creates a new image from a data stream. The name parameter is required to uniquely identify this image (for caching etc). If the image in this io.Reader is an SVG, the name should end ".svg". Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewImageFromResource

```go
func NewImageFromResource(res fyne.Resource) *Image
```
NewImageFromResource creates a new image by loading the specified resource. Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.

#### func  NewImageFromURI

```go
func NewImageFromURI(uri fyne.URI) *Image
```
NewImageFromURI creates a new image from named resource. File URIs will read the file path and other schemes will download the data into a resource. HTTP and HTTPs URIs will use the GET method by default to request the resource. Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Image) Alpha

```go
func (i *Image) Alpha() float64
```
Alpha is a convenience function that returns the alpha value for an image based on its Translucency value. The result is 1.0 - Translucency.

#### func (*Image) Hide

```go
func (o *Image) Hide()
```
Hide will set this object to not be visible.

#### func (*Image) MinSize

```go
func (o *Image) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise.

#### func (*Image) Move

```go
func (o *Image) Move(pos fyne.Position)
```
Move the object to a new position, relative to its parent.

#### func (*Image) Position

```go
func (o *Image) Position() fyne.Position
```
Position gets the current position of this canvas object, relative to its parent.

#### func (*Image) Refresh

```go
func (i *Image) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Image) Resize

```go
func (i *Image) Resize(s fyne.Size)
```
Resize on an image will scale the content or reposition it according to FillMode. It will normally cause a Refresh to ensure the pixels are recalculated.

#### func (*Image) SetMinSize

```go
func (o *Image) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be.

#### func (*Image) Show

```go
func (o *Image) Show()
```
Show will set this object to be visible.

#### func (*Image) Size

```go
func (o *Image) Size() fyne.Size
```
Size returns the current size of this canvas object.

#### func (*Image) Visible

```go
func (o *Image) Visible() bool
```
Visible returns true if this object is visible, false otherwise.
