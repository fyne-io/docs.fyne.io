---
tags: [api]
title: widget.ImageSegment
slug: imagesegment

aliases:
- /api/v2.7/widget/imagesegment

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type ImageSegment struct {
	Source fyne.URI
	Title  string

	// Alignment specifies the horizontal alignment of this image segment
	// Since: 2.4
	Alignment fyne.TextAlign
}
```

ImageSegment represents an image within a rich text widget.


<div class="since">Since: <code>
2.3</code></div>

###

```go
func (i *ImageSegment) Inline() bool
```
Inline returns false as images in rich text are blocks.

###

```go
func (i *ImageSegment) Select(begin, end fyne.Position)
```
Select tells the segment that the user is selecting the content between the two positions.

###

```go
func (i *ImageSegment) SelectedText() string
```
SelectedText should return the text representation of any content currently selected through the Select call.

###

```go
func (i *ImageSegment) Textual() string
```
Textual returns the content of this segment rendered to plain text.

###

```go
func (i *ImageSegment) Unselect()
```
Unselect tells the segment that the user is has cancelled the previous selection.

###

```go
func (i *ImageSegment) Update(o fyne.CanvasObject)
```
Update applies the current state of this image segment to an existing visual.

###

```go
func (i *ImageSegment) Visual() fyne.CanvasObject
```
Visual returns a new instance of an image widget required to render this segment.
