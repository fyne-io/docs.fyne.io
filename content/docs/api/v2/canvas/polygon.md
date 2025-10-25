---
tags: [api]
title: canvas.Polygon
slug: polygon

aliases:
- /api/v2/canvas/polygon.html
- /api/v2.0/canvas/polygon
- /api/v2.0/canvas/polygon.html
- /api/v2.1/canvas/polygon
- /api/v2.1/canvas/polygon.html
- /api/v2.2/canvas/polygon
- /api/v2.2/canvas/polygon.html
- /api/v2.3/canvas/polygon
- /api/v2.3/canvas/polygon.html
- /api/v2.4/canvas/polygon
- /api/v2.4/canvas/polygon.html
- /api/v2.5/canvas/polygon
- /api/v2.5/canvas/polygon.html
- /api/v2.6/canvas/polygon
- /api/v2.6/canvas/polygon.html
- /api/v2.7/canvas/polygon
- /api/v2.7/canvas/polygon.html

package: fyne.io/fyne/v2/canvas
---


---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type Polygon

```go
type Polygon struct {
	FillColor    color.Color // The polygon fill color
	StrokeColor  color.Color // The polygon stroke color
	StrokeWidth  float32     // The stroke width of the polygon
	CornerRadius float32     // The radius of the polygon corners
	Angle        float32     // Angle of polygon, in degrees (positive means clockwise, negative means counter-clockwise direction).
	Sides        uint        //	Amount of sides of polygon.
}
```

Polygon describes a colored regular polygon primitive in a Fyne canvas. The rendered portion will be in the center of the available space.


<div class="since">Since: <code>
2.7</code></div>

#### func  NewPolygon

```go
func NewPolygon(sides uint, color color.Color) *Polygon
```
NewPolygon returns a new Polygon instance

#### func (*Polygon) Hide

```go
func (r *Polygon) Hide()
```
Hide will set this polygon to not be visible

#### func (*Polygon) MinSize

```go
func (o *Polygon) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise.

#### func (*Polygon) Move

```go
func (r *Polygon) Move(pos fyne.Position)
```
Move the polygon to a new position, relative to its parent / canvas

#### func (*Polygon) Position

```go
func (o *Polygon) Position() fyne.Position
```
Position gets the current position of this canvas object, relative to its parent.

#### func (*Polygon) Refresh

```go
func (r *Polygon) Refresh()
```
Refresh causes this polygon to be redrawn with its configured state.

#### func (*Polygon) Resize

```go
func (r *Polygon) Resize(s fyne.Size)
```
Resize on a polygon updates the new size of this object. If it has a stroke width this will cause it to Refresh.

#### func (*Polygon) SetMinSize

```go
func (o *Polygon) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be.

#### func (*Polygon) Show

```go
func (o *Polygon) Show()
```
Show will set this object to be visible.

#### func (*Polygon) Size

```go
func (o *Polygon) Size() fyne.Size
```
Size returns the current size of this canvas object.

#### func (*Polygon) Visible

```go
func (o *Polygon) Visible() bool
```
Visible returns true if this object is visible, false otherwise.
