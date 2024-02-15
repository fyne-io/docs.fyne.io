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

#### type Rectangle

```go
type Rectangle struct {
	FillColor   color.Color // The rectangle fill colour
	StrokeColor color.Color // The rectangle stroke colour
	StrokeWidth float32     // The stroke width of the rectangle
}
```

Rectangle describes a coloured rectangle primitive in a Fyne canvas

#### func  NewRectangle

```go
func NewRectangle(color color.Color) *Rectangle
```
NewRectangle returns a new Rectangle instance

#### func (*Rectangle) Hide

```go
func (r *Rectangle) Hide()
```
Hide will set this object to not be visible

#### func (*Rectangle) MinSize

```go
func (r *Rectangle) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise

#### func (*Rectangle) Move

```go
func (r *Rectangle) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*Rectangle) Position

```go
func (r *Rectangle) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to its parent / canvas

#### func (*Rectangle) Refresh

```go
func (r *Rectangle) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Rectangle) Resize

```go
func (r *Rectangle) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*Rectangle) SetMinSize

```go
func (r *Rectangle) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*Rectangle) Show

```go
func (r *Rectangle) Show()
```
Show will set this object to be visible

#### func (*Rectangle) Size

```go
func (r *Rectangle) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*Rectangle) Visible

```go
func (r *Rectangle) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise
