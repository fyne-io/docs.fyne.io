---
layout: page
tags: [api]
title: Fyne API "widget.TextSegment"
package: fyne.io/fyne/v2/widget
---

# widget.TextSegment
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type TextSegment

```go
type TextSegment struct {
	Style RichTextStyle
	Text  string
}
```

TextSegment represents the styling for a segment of rich text.


<div class="since">Since: <code>
2.1</code></div>

#### func (*TextSegment) Inline

```go
func (t *TextSegment) Inline() bool
```
Inline should return true if this text can be included within other elements, or false if it creates a new block.

#### func (*TextSegment) Select

```go
func (t *TextSegment) Select(begin, end fyne.Position)
```
Select tells the segment that the user is selecting the content between the two positions.

#### func (*TextSegment) SelectedText

```go
func (t *TextSegment) SelectedText() string
```
SelectedText should return the text representation of any content currently selected through the Select call.

#### func (*TextSegment) Textual

```go
func (t *TextSegment) Textual() string
```
Textual returns the content of this segment rendered to plain text.

#### func (*TextSegment) Unselect

```go
func (t *TextSegment) Unselect()
```
Unselect tells the segment that the user is has cancelled the previous selection.

#### func (*TextSegment) Update

```go
func (t *TextSegment) Update(o fyne.CanvasObject)
```
Update applies the current state of this text segment to an existing visual.

#### func (*TextSegment) Visual

```go
func (t *TextSegment) Visual() fyne.CanvasObject
```
Visual returns the graphical elements required to render this segment.
