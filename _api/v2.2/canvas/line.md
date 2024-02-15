---
layout: page
tags: [api]
title: Fyne API "canvas.Line"
---

# canvas.Line
---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type Line

```go
type Line struct {
	Position1 fyne.Position // The current top-left position of the Line
	Position2 fyne.Position // The current bottomright position of the Line
	Hidden    bool          // Is this Line currently hidden

	StrokeColor color.Color // The line stroke color
	StrokeWidth float32     // The stroke width of the line
}
```

Line describes a colored line primitive in a Fyne canvas. Lines are special as they can have a negative width or height to indicate an inverse slope (i.e. slope up vs down).

#### func  NewLine

```go
func NewLine(color color.Color) *Line
```
NewLine returns a new Line instance

#### func (*Line) Hide

```go
func (l *Line) Hide()
```
Hide will set this line to not be visible

#### func (*Line) MinSize

```go
func (l *Line) MinSize() fyne.Size
```
MinSize for a Line simply returns Size{1, 1} as there is no explicit content

#### func (*Line) Move

```go
func (l *Line) Move(pos fyne.Position)
```
Move the line object to a new position, relative to its parent / canvas

#### func (*Line) Position

```go
func (l *Line) Position() fyne.Position
```
Position gets the current top-left position of this line object, relative to its parent / canvas

#### func (*Line) Refresh

```go
func (l *Line) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Line) Resize

```go
func (l *Line) Resize(size fyne.Size)
```
Resize sets a new bottom-right position for the line object and it will then be refreshed.

#### func (*Line) Show

```go
func (l *Line) Show()
```
Show will set this line to be visible

#### func (*Line) Size

```go
func (l *Line) Size() fyne.Size
```
Size returns the current size of bounding box for this line object

#### func (*Line) Visible

```go
func (l *Line) Visible() bool
```
Visible returns true if this line// Show will set this circle to be visible is visible, false otherwise
