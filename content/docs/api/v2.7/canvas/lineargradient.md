---
tags: [api]
title: canvas.LinearGradient
slug: lineargradient

aliases:
- /api/v2.7/canvas/lineargradient

package: fyne.io/fyne/v2/canvas
---


---
```go
import "fyne.io/fyne/v2/canvas"
```

#

###

```go
type LinearGradient struct {
	StartColor color.Color // The beginning color of the gradient
	EndColor   color.Color // The end color of the gradient
	Angle      float64     // The angle of the gradient (0/180 for vertical; 90/270 for horizontal)
}
```

LinearGradient defines a Gradient travelling straight at a given angle. The only supported values for the angle are `0.0` (vertical) and `90.0` (horizontal), currently.

###

```go
func NewHorizontalGradient(start, end color.Color) *LinearGradient
```
NewHorizontalGradient creates a new horizontally travelling linear gradient. The start color will be at the left of the gradient and the end color will be at the right.

###

```go
func NewLinearGradient(start, end color.Color, angle float64) *LinearGradient
```
NewLinearGradient creates a linear gradient at the specified angle. The angle parameter is the degree angle along which the gradient is calculated. A NewHorizontalGradient uses 270 degrees and NewVerticalGradient is 0 degrees.

###

```go
func NewVerticalGradient(start color.Color, end color.Color) *LinearGradient
```
NewVerticalGradient creates a new vertically travelling linear gradient. The start color will be at the top of the gradient and the end color will be at the bottom.

###

```go
func (g *LinearGradient) Generate(iw, ih int) image.Image
```
Generate calculates an image of the gradient with the specified width and height.

###

```go
func (g *LinearGradient) Hide()
```
Hide will set this gradient to not be visible

###

```go
func (o *LinearGradient) MinSize() fyne.Size
```
MinSize returns the specified minimum size, if set, or {1, 1} otherwise.

###

```go
func (g *LinearGradient) Move(pos fyne.Position)
```
Move the gradient to a new position, relative to its parent / canvas

###

```go
func (o *LinearGradient) Position() fyne.Position
```
Position gets the current position of this canvas object, relative to its parent.

###

```go
func (g *LinearGradient) Refresh()
```
Refresh causes this gradient to be redrawn with its configured state.

###

```go
func (g *LinearGradient) Resize(size fyne.Size)
```
Resize resizes the gradient to a new size.

###

```go
func (o *LinearGradient) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be.

###

```go
func (o *LinearGradient) Show()
```
Show will set this object to be visible.

###

```go
func (o *LinearGradient) Size() fyne.Size
```
Size returns the current size of this canvas object.

###

```go
func (o *LinearGradient) Visible() bool
```
Visible returns true if this object is visible, false otherwise.
