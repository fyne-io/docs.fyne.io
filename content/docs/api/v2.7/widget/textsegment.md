---
tags: [api]
title: widget.TextSegment
slug: textsegment

aliases:
- /api/v2.7/widget/textsegment

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type TextSegment struct {
	Style RichTextStyle
	Text  string
}
```

TextSegment represents the styling for a segment of rich text.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (t *TextSegment) Inline() bool
```
Inline should return true if this text can be included within other elements, or false if it creates a new block.

###

```go
func (t *TextSegment) Select(begin, end fyne.Position)
```
Select tells the segment that the user is selecting the content between the two positions.

###

```go
func (t *TextSegment) SelectedText() string
```
SelectedText should return the text representation of any content currently selected through the Select call.

###

```go
func (t *TextSegment) Textual() string
```
Textual returns the content of this segment rendered to plain text.

###

```go
func (t *TextSegment) Unselect()
```
Unselect tells the segment that the user is has cancelled the previous selection.

###

```go
func (t *TextSegment) Update(o fyne.CanvasObject)
```
Update applies the current state of this text segment to an existing visual.

###

```go
func (t *TextSegment) Visual() fyne.CanvasObject
```
Visual returns a new instance of a graphical element required to render this segment.
