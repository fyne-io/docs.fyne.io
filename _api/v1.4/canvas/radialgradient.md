---
layout: page
tags: [api]
title: Fyne API "canvas.RadialGradient"
---

# canvas.RadialGradient
---
```go
import "fyne.io/fyne/canvas"
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
func (r *RadialGradient) Hide()
```
Hide will set this object to not be visible

#### func (*RadialGradient) MinSize

```go
func (r *RadialGradient) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise

#### func (*RadialGradient) Move

```go
func (r *RadialGradient) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*RadialGradient) Position

```go
func (r *RadialGradient) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to its parent / canvas

#### func (*RadialGradient) Refresh

```go
func (g *RadialGradient) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*RadialGradient) Resize

```go
func (r *RadialGradient) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*RadialGradient) SetMinSize

```go
func (r *RadialGradient) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*RadialGradient) Show

```go
func (r *RadialGradient) Show()
```
Show will set this object to be visible

#### func (*RadialGradient) Size

```go
func (r *RadialGradient) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*RadialGradient) Visible

```go
func (r *RadialGradient) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise
