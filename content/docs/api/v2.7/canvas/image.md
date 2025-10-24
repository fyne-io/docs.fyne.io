---
tags: [api]
title: canvas.Image
slug: image

aliases:
- /api/v2.7/canvas/image

package: fyne.io/fyne/v2/canvas
---


---
```go
import "fyne.io/fyne/v2/canvas"
```

#

###

```go
type Image struct {

	// one of the following sources will provide our image data
	File     string        // Load the image from a file
	Resource fyne.Resource // Load the image from an in-memory resource
	Image    image.Image   // Specify a loaded image to use in this canvas object

	Translucency float64    // Set a translucency value > 0.0 to fade the image
	FillMode     ImageFill  // Specify how the image should expand to fill or fit the available space
	ScaleMode    ImageScale // Specify the type of scaling interpolation applied to the image

	// CornerRadius specifies a radius to apply to round corners of the image.
	//
	// Since: 2.7
	CornerRadius float32
}
```

Image describes a drawable image area that can render in a Fyne canvas The image may be a vector or a bitmap representation, it will fill the area. The fill mode can be changed by setting FillMode to a different ImageFill.

###

```go
func NewImageFromFile(file string) *Image
```
NewImageFromFile creates a new image from a local file. Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.

###

```go
func NewImageFromImage(img image.Image) *Image
```
NewImageFromImage returns a new Image instance that is rendered from the Go image.Image passed in. Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.

###

```go
func NewImageFromReader(read io.Reader, name string) *Image
```
NewImageFromReader creates a new image from a data stream. The name parameter is required to uniquely identify this image (for caching etc.). If the image in this io.Reader is an SVG, the name should end ".svg". Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func NewImageFromResource(res fyne.Resource) *Image
```
NewImageFromResource creates a new image by loading the specified resource. Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.

###

```go
func NewImageFromURI(uri fyne.URI) *Image
```
NewImageFromURI creates a new image from named resource. File URIs will read the file path and other schemes will download the data into a resource. HTTP and HTTPs URIs will use the GET method by default to request the resource. Images returned from this method will scale to fit the canvas object. The method for scaling can be set using the Fill field.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func (i *Image) Alpha() float64
```
Alpha is a convenience function that returns the alpha value for an image based on its Translucency value. The result is 1.0 - Translucency.

###

```go
func (i *Image) Aspect() float32
```
Aspect will return the original content aspect after it was last refreshed.


<div class="since">Since: <code>
2.4</code></div>

###

```go
func (i *Image) Hide()
```
Hide will set this image to not be visible

###

```go
func (i *Image) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise.

###

```go
func (i *Image) Move(pos fyne.Position)
```
Move the image object to a new position, relative to its parent top, left corner.

###

```go
func (o *Image) Position() fyne.Position
```
Position gets the current position of this canvas object, relative to its parent.

###

```go
func (i *Image) Refresh()
```
Refresh causes this image to be redrawn with its configured state.

###

```go
func (i *Image) Resize(s fyne.Size)
```
Resize on an image will scale the content or reposition it according to FillMode. It will normally cause a Refresh to ensure the pixels are recalculated.

###

```go
func (o *Image) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be.

###

```go
func (o *Image) Show()
```
Show will set this object to be visible.

###

```go
func (o *Image) Size() fyne.Size
```
Size returns the current size of this canvas object.

###

```go
func (o *Image) Visible() bool
```
Visible returns true if this object is visible, false otherwise.
