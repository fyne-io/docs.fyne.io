---
layout: page
tags: [api]
title: Fyne API "canvas.Rectangle"
package: fyne.io/fyne/v2/canvas
---

# canvas.Rectangle
---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type Rectangle

```go
type Rectangle struct {
	FillColor   color.Color // The rectangle fill color
	StrokeColor color.Color // The rectangle stroke color
	StrokeWidth float32     // The stroke width of the rectangle
	// The radius of the rectangle corners
	//
	// Since: 2.4
	CornerRadius float32
}
```

Rectangle describes a colored rectangle primitive in a Fyne canvas

#### func  NewRectangle

```go
func NewRectangle(color color.Color) *Rectangle
```
NewRectangle returns a new Rectangle instance

#### func (*Rectangle) Hide

```go
func (r *Rectangle) Hide()
```
Hide will set this rectangle to not be visible

#### func (*Rectangle) MinSize

```go
func (o *Rectangle) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise.

#### func (*Rectangle) Move

```go
func (r *Rectangle) Move(pos fyne.Position)
```
Move the rectangle to a new position, relative to its parent / canvas

#### func (*Rectangle) Position

```go
func (o *Rectangle) Position() fyne.Position
```
Position gets the current position of this canvas object, relative to its parent.

#### func (*Rectangle) Refresh

```go
func (r *Rectangle) Refresh()
```
Refresh causes this rectangle to be redrawn with its configured state.

#### func (*Rectangle) Resize

```go
func (r *Rectangle) Resize(s fyne.Size)
```
Resize on a rectangle updates the new size of this object. If it has a stroke width this will cause it to Refresh.

#### func (*Rectangle) SetMinSize

```go
func (o *Rectangle) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be.

#### func (*Rectangle) Show

```go
func (o *Rectangle) Show()
```
Show will set this object to be visible.

#### func (*Rectangle) Size

```go
func (o *Rectangle) Size() fyne.Size
```
Size returns the current size of this canvas object.

#### func (*Rectangle) Visible

```go
func (o *Rectangle) Visible() bool
```
Visible returns true if this object is visible, false otherwise.
