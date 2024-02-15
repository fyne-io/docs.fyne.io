---
layout: page
tags: [api]
title: Fyne API canvas
---

# canvas
---
```go
import "fyne.io/fyne/canvas"
```

## Usage

#### type Text

```go
type Text struct {
	Alignment fyne.TextAlign // The alignment of the text content

	Color     color.Color    // The main text draw colour
	Text      string         // The string content of this Text
	TextSize  int            // Size of the text - if the Canvas scale is 1.0 this will be equivalent to point size
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
Hide will set this object to not be visible

#### func (*Text) MinSize

```go
func (t *Text) MinSize() fyne.Size
```
MinSize returns the minimum size of this text object based on its font size and content. This is normally determined by the render implementation.

#### func (*Text) Move

```go
func (r *Text) Move(pos fyne.Position)
```
Move the rectangle object to a new position, relative to its parent / canvas

#### func (*Text) Position

```go
func (r *Text) Position() fyne.Position
```
CurrentPosition gets the current position of this rectangle object, relative to its parent / canvas

#### func (*Text) Refresh

```go
func (t *Text) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Text) Resize

```go
func (r *Text) Resize(size fyne.Size)
```
Resize sets a new size for the rectangle object

#### func (*Text) SetMinSize

```go
func (r *Text) SetMinSize(size fyne.Size)
```
SetMinSize specifies the smallest size this object should be

#### func (*Text) Show

```go
func (r *Text) Show()
```
Show will set this object to be visible

#### func (*Text) Size

```go
func (r *Text) Size() fyne.Size
```
CurrentSize returns the current size of this rectangle object

#### func (*Text) Visible

```go
func (r *Text) Visible() bool
```
IsVisible returns true if this object is visible, false otherwise
