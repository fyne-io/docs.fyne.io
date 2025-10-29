---
tags: [api]
title: canvas.RadialGradient
slug: radialgradient

aliases:
- /api/canvas/radialgradient
- /api/canvas/radialgradient.html
- /api/v2.0/canvas/radialgradient
- /api/v2.0/canvas/radialgradient.html
- /api/v2.1/canvas/radialgradient
- /api/v2.1/canvas/radialgradient.html
- /api/v2.2/canvas/radialgradient
- /api/v2.2/canvas/radialgradient.html
- /api/v2.3/canvas/radialgradient
- /api/v2.3/canvas/radialgradient.html
- /api/v2.4/canvas/radialgradient
- /api/v2.4/canvas/radialgradient.html
- /api/v2.5/canvas/radialgradient
- /api/v2.5/canvas/radialgradient.html
- /api/v2.6/canvas/radialgradient
- /api/v2.6/canvas/radialgradient.html
- /api/v2.7/canvas/radialgradient
- /api/v2.7/canvas/radialgradient.html

package: fyne.io/fyne/v2/canvas
---


---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type RadialGradient

```go
type RadialGradient struct {
	StartColor color.Color // The beginning color of the gradient
	EndColor   color.Color // The end color of the gradient
	// The offset of the center for generation of the gradient.
	// This is not a DP measure but relates to the width/height.
	// A value of 0.5 would move the center by the half width/height.
	CenterOffsetX, CenterOffsetY float64
}
```

RadialGradient defines a Gradient travelling radially from a center point outward.

#### func  NewRadialGradient

```go
func NewRadialGradient(start, end color.Color) *RadialGradient
```
NewRadialGradient creates a new radial gradient.

#### func (*RadialGradient) Generate

```go
func (g *RadialGradient) Generate(iw, ih int) image.Image
```
Generate calculates an image of the gradient with the specified width and height.

#### func (*RadialGradient) Hide

```go
func (g *RadialGradient) Hide()
```
Hide will set this gradient to not be visible

#### func (*RadialGradient) MinSize

```go
func (o *RadialGradient) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise.

#### func (*RadialGradient) Move

```go
func (g *RadialGradient) Move(pos fyne.Position)
```
Move the gradient to a new position, relative to its parent / canvas

#### func (*RadialGradient) Position

```go
func (o *RadialGradient) Position() fyne.Position
```
Position gets the current position of this canvas object, relative to its parent.

#### func (*RadialGradient) Refresh

```go
func (g *RadialGradient) Refresh()
```
Refresh causes this gradient to be redrawn with its configured state.

#### func (*RadialGradient) Resize

```go
func (g *RadialGradient) Resize(size fyne.Size)
```
Resize resizes the gradient to a new size.

#### func (*RadialGradient) SetMinSize

```go
func (o *RadialGradient) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be.

#### func (*RadialGradient) Show

```go
func (o *RadialGradient) Show()
```
Show will set this object to be visible.

#### func (*RadialGradient) Size

```go
func (o *RadialGradient) Size() fyne.Size
```
Size returns the current size of this canvas object.

#### func (*RadialGradient) Visible

```go
func (o *RadialGradient) Visible() bool
```
Visible returns true if this object is visible, false otherwise.
