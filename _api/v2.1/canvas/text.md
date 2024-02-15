---
layout: page
tags: [api]
title: Fyne API "canvas.Text"
---

# canvas.Text
---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type Text

```go
type Text struct {
	Alignment fyne.TextAlign // The alignment of the text content

	Color     color.Color    // The main text draw color
	Text      string         // The string content of this Text
	TextSize  float32        // Size of the text - if the Canvas scale is 1.0 this will be equivalent to point size
	TextStyle fyne.TextStyle // The style of the text content
}
```

Text describes a text primitive in a Fyne canvas. A text object can have a style set which will apply to the whole string. No formatting or text parsing will be performed

#### func  NewText

```go
func NewText(text string, color color.Color) *Text
```
NewText returns a new Text implementation

#### func (*Text) Hide

```go
func (r *Text) Hide()
```
Hide will set this object to not be visible.

#### func (*Text) MinSize

```go
func (t *Text) MinSize() fyne.Size
```
MinSize returns the minimum size of this text object based on its font size and content. This is normally determined by the render implementation.

#### func (*Text) Move

```go
func (r *Text) Move(pos fyne.Position)
```
Move the object to a new position, relative to its parent.

#### func (*Text) Position

```go
func (r *Text) Position() fyne.Position
```
CurrentPosition gets the current position of this canvas object, relative to its parent.

#### func (*Text) Refresh

```go
func (t *Text) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Text) Resize

```go
func (r *Text) Resize(size fyne.Size)
```
Resize sets a new size for the canvas object.

#### func (*Text) SetMinSize

```go
func (t *Text) SetMinSize(size fyne.Size)
```
SetMinSize has no effect as the smallest size this canvas object can be is based on its font size and content.

#### func (*Text) Show

```go
func (r *Text) Show()
```
Show will set this object to be visible.

#### func (*Text) Size

```go
func (r *Text) Size() fyne.Size
```
CurrentSize returns the current size of this canvas object.

#### func (*Text) Visible

```go
func (r *Text) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise.
