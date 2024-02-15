---
layout: page
tags: [api]
title: Fyne API "canvas.LinearGradient"
---

# canvas.LinearGradient
---
```go
import "fyne.io/fyne/canvas"
```

## Usage

#### type LinearGradient

```go
type LinearGradient struct {
	StartColor color.Color // The beginning color of the gradient
	EndColor   color.Color // The end color of the gradient
	Angle      float64     // The angle of the gradient (0/180 for vertical; 90/270 for horizontal)
}
```

LinearGradient defines a Gradient travelling straight at a given angle. The only supported values for the angle are `0.0` (vertical) and `90.0` (horizontal), currently.

#### func  NewHorizontalGradient

```go
func NewHorizontalGradient(start, end color.Color) *LinearGradient
```
NewHorizontalGradient creates a new horizontally travelling linear gradient. The start color will be at the left of the gradient and the end color will be at the right.

#### func  NewLinearGradient

```go
func NewLinearGradient(start, end color.Color, angle float64) *LinearGradient
```
NewLinearGradient creates a linear gradient at a the specified angle. The angle parameter is the degree angle along which the gradient is calculated. A NewHorizontalGradient uses 270 degrees and NewVerticalGradient is 0 degrees.

#### func  NewVerticalGradient

```go
func NewVerticalGradient(start color.Color, end color.Color) *LinearGradient
```
NewVerticalGradient creates a new vertically travelling linear gradient. The start color will be at the top of the gradient and the end color will be at the bottom.

#### func (*LinearGradient) Generate

```go
func (g *LinearGradient) Generate(iw, ih int) image.Image
```
Generate calculates an image of the gradient with the specified width and height.

#### func (*LinearGradient) Hide

```go
func (r *LinearGradient) Hide()
```
Hide will set this object to not be visible

#### func (*LinearGradient) MinSize

```go
func (r *LinearGradient) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise

#### func (*LinearGradient) Move

```go
func (r *LinearGradient) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*LinearGradient) Position

```go
func (r *LinearGradient) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to its parent / canvas

#### func (*LinearGradient) Refresh

```go
func (g *LinearGradient) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*LinearGradient) Resize

```go
func (r *LinearGradient) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*LinearGradient) SetMinSize

```go
func (r *LinearGradient) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*LinearGradient) Show

```go
func (r *LinearGradient) Show()
```
Show will set this object to be visible

#### func (*LinearGradient) Size

```go
func (r *LinearGradient) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*LinearGradient) Visible

```go
func (r *LinearGradient) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise
