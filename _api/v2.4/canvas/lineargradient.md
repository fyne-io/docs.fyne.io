---
layout: page
tags: [api]
title: Fyne API "canvas.LinearGradient"
package: fyne.io/fyne/v2/canvas
---

# canvas.LinearGradient
---
```go
import "fyne.io/fyne/v2/canvas"
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
NewLinearGradient creates a linear gradient at the specified angle. The angle parameter is the degree angle along which the gradient is calculated. A NewHorizontalGradient uses 270 degrees and NewVerticalGradient is 0 degrees.

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
func (g *LinearGradient) Hide()
```
Hide will set this gradient to not be visible

#### func (*LinearGradient) MinSize

```go
func (o *LinearGradient) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise.

#### func (*LinearGradient) Move

```go
func (g *LinearGradient) Move(pos fyne.Position)
```
Move the gradient to a new position, relative to its parent / canvas

#### func (*LinearGradient) Position

```go
func (o *LinearGradient) Position() fyne.Position
```
Position gets the current position of this canvas object, relative to its parent.

#### func (*LinearGradient) Refresh

```go
func (g *LinearGradient) Refresh()
```
Refresh causes this gradient to be redrawn with its configured state.

#### func (*LinearGradient) Resize

```go
func (o *LinearGradient) Resize(size fyne.Size)
```
Resize sets a new size for the canvas object.

#### func (*LinearGradient) SetMinSize

```go
func (o *LinearGradient) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be.

#### func (*LinearGradient) Show

```go
func (o *LinearGradient) Show()
```
Show will set this object to be visible.

#### func (*LinearGradient) Size

```go
func (o *LinearGradient) Size() fyne.Size
```
Size returns the current size of this canvas object.

#### func (*LinearGradient) Visible

```go
func (o *LinearGradient) Visible() bool
```
Visible returns true if this object is visible, false otherwise.
