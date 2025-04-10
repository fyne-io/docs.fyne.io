---
layout: page
tags: [api]
title: Fyne API "canvas.Circle"
package: fyne.io/fyne/v2/canvas
---

# canvas.Circle
---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type Circle

```go
type Circle struct {
	Position1 fyne.Position // The current top-left position of the Circle
	Position2 fyne.Position // The current bottomright position of the Circle
	Hidden    bool          // Is this circle currently hidden

	FillColor   color.Color // The circle fill color
	StrokeColor color.Color // The circle stroke color
	StrokeWidth float32     // The stroke width of the circle
}
```

Circle describes a colored circle primitive in a Fyne canvas

#### func  NewCircle

```go
func NewCircle(color color.Color) *Circle
```
NewCircle returns a new Circle instance

#### func (*Circle) Hide

```go
func (c *Circle) Hide()
```
Hide will set this circle to not be visible

#### func (*Circle) MinSize

```go
func (c *Circle) MinSize() fyne.Size
```
MinSize for a Circle simply returns Size{1, 1} as there is no explicit content

#### func (*Circle) Move

```go
func (c *Circle) Move(pos fyne.Position)
```
Move the circle object to a new position, relative to its parent / canvas

#### func (*Circle) Position

```go
func (c *Circle) Position() fyne.Position
```
Position gets the current top-left position of this circle object, relative to its parent / canvas

#### func (*Circle) Refresh

```go
func (c *Circle) Refresh()
```
Refresh causes this object to be redrawn with its configured state.

#### func (*Circle) Resize

```go
func (c *Circle) Resize(size fyne.Size)
```
Resize sets a new bottom-right position for the circle object If it has a stroke width this will cause it to Refresh.

#### func (*Circle) Show

```go
func (c *Circle) Show()
```
Show will set this circle to be visible

#### func (*Circle) Size

```go
func (c *Circle) Size() fyne.Size
```
Size returns the current size of bounding box for this circle object

#### func (*Circle) Visible

```go
func (c *Circle) Visible() bool
```
Visible returns true if this circle is visible, false otherwise
