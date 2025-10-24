---
tags: [api]
title: canvas.Rectangle
slug: rectangle

aliases:
- /api/v2.7/canvas/rectangle

package: fyne.io/fyne/v2/canvas
---


---
```go
import "fyne.io/fyne/v2/canvas"
```

#

###

```go
type Rectangle struct {
	FillColor   color.Color // The rectangle fill color
	StrokeColor color.Color // The rectangle stroke color
	StrokeWidth float32     // The stroke width of the rectangle
	// The radius of the rectangle corners
	//
	// Since: 2.4
	CornerRadius float32

	// Enforce an aspect ratio for the rectangle, the content will be made shorter or narrower
	// to meet the requested aspect, if set.
	//
	// Since: 2.7
	Aspect float32

	// The radius of the rectangle top-right corner only.
	//
	// Since: 2.7
	TopRightCornerRadius float32

	// The radius of the rectangle top-left corner only.
	//
	// Since: 2.7
	TopLeftCornerRadius float32

	// The radius of the rectangle bottom-right corner only.
	//
	// Since: 2.7
	BottomRightCornerRadius float32

	// The radius of the rectangle bottom-left corner only.
	//
	// Since: 2.7
	BottomLeftCornerRadius float32
}
```

Rectangle describes a colored rectangle primitive in a Fyne canvas

###

```go
func NewRectangle(color color.Color) *Rectangle
```
NewRectangle returns a new Rectangle instance

###

```go
func NewSquare(color color.Color) *Rectangle
```
NewSquare returns a new Rectangle instance that has a square aspect ratio.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func (r *Rectangle) Hide()
```
Hide will set this rectangle to not be visible

###

```go
func (o *Rectangle) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise.

###

```go
func (r *Rectangle) Move(pos fyne.Position)
```
Move the rectangle to a new position, relative to its parent / canvas

###

```go
func (o *Rectangle) Position() fyne.Position
```
Position gets the current position of this canvas object, relative to its parent.

###

```go
func (r *Rectangle) Refresh()
```
Refresh causes this rectangle to be redrawn with its configured state.

###

```go
func (r *Rectangle) Resize(s fyne.Size)
```
Resize on a rectangle updates the new size of this object. If it has a stroke width this will cause it to Refresh. If Aspect is non-zero it may cause the rectangle to be smaller than the requested size.

###

```go
func (o *Rectangle) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be.

###

```go
func (o *Rectangle) Show()
```
Show will set this object to be visible.

###

```go
func (o *Rectangle) Size() fyne.Size
```
Size returns the current size of this canvas object.

###

```go
func (o *Rectangle) Visible() bool
```
Visible returns true if this object is visible, false otherwise.
