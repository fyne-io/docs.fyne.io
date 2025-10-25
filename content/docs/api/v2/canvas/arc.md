---
tags: [api]
title: canvas.Arc
slug: arc

aliases:
- /api/v2/canvas/arc.html
- /api/v2.0/canvas/arc
- /api/v2.0/canvas/arc.html
- /api/v2.1/canvas/arc
- /api/v2.1/canvas/arc.html
- /api/v2.2/canvas/arc
- /api/v2.2/canvas/arc.html
- /api/v2.3/canvas/arc
- /api/v2.3/canvas/arc.html
- /api/v2.4/canvas/arc
- /api/v2.4/canvas/arc.html
- /api/v2.5/canvas/arc
- /api/v2.5/canvas/arc.html
- /api/v2.6/canvas/arc
- /api/v2.6/canvas/arc.html
- /api/v2.7/canvas/arc
- /api/v2.7/canvas/arc.html

package: fyne.io/fyne/v2/canvas
---


---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type Arc

```go
type Arc struct {
	FillColor    color.Color // The arc fill colour
	StartAngle   float32     // Start angle in degrees
	EndAngle     float32     // End angle in degrees
	CornerRadius float32     // Radius used to round the corners
	StrokeColor  color.Color // The arc stroke color
	StrokeWidth  float32     // The stroke width of the arc
	CutoutRatio  float32     // Controls what portion of the inner should be cut out. A value of 0.0 results in a pie slice, while 1.0 results in a stroke.
}
```

Arc represents a filled arc or annular sector primitive that can be drawn on a Fyne canvas. It allows for the creation of circular, ring-shaped or pie-shaped segment, with configurable cutout ratio as well as customizable start and end angles to define the arc's length as the absolute difference between the two angles. The arc is always centered on its position. The arc is drawn from StartAngle to EndAngle (in degrees, positive is clockwise, negative is counter-clockwise). 0°/360 is top, 90° is right, 180° is bottom, 270° is left 0°/-360 is top, -90° is left, -180° is bottom, -270° is right


<div class="since">Since: <code>
2.7</code></div>

#### func  NewArc

```go
func NewArc(startAngle, endAngle, cutoutRatio float32, color color.Color) *Arc
```
NewArc returns a new Arc instance with the specified start and end angles (in degrees), fill color and cutout ratio.

#### func  NewDoughnutArc

```go
func NewDoughnutArc(startAngle, endAngle float32, color color.Color) *Arc
```
NewDoughnutArc returns a new doughnut-shaped Arc instance with the specified start and end angles (in degrees), fill color and cutout ratio set to 0.5.

#### func  NewPieArc

```go
func NewPieArc(startAngle, endAngle float32, color color.Color) *Arc
```
NewPieArc returns a new pie-shaped Arc instance with the specified start and end angles (in degrees), fill color and cutout ratio set to 0.

#### func (*Arc) Hide

```go
func (a *Arc) Hide()
```
Hide will set this arc to not be visible.

#### func (*Arc) MinSize

```go
func (o *Arc) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise.

#### func (*Arc) Move

```go
func (a *Arc) Move(pos fyne.Position)
```
Move the arc to a new position, relative to its parent / canvas. The position specifies the **center** of the arc.

#### func (*Arc) Position

```go
func (o *Arc) Position() fyne.Position
```
Position gets the current position of this canvas object, relative to its parent.

#### func (*Arc) Refresh

```go
func (a *Arc) Refresh()
```
Refresh causes this arc to be redrawn with its configured state.

#### func (*Arc) Resize

```go
func (a *Arc) Resize(s fyne.Size)
```
Resize updates the logical size of the arc. The arc is always drawn centered on its Position().

#### func (*Arc) SetMinSize

```go
func (o *Arc) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be.

#### func (*Arc) Show

```go
func (o *Arc) Show()
```
Show will set this object to be visible.

#### func (*Arc) Size

```go
func (o *Arc) Size() fyne.Size
```
Size returns the current size of this canvas object.

#### func (*Arc) Visible

```go
func (o *Arc) Visible() bool
```
Visible returns true if this object is visible, false otherwise.
