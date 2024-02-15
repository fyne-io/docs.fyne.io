---
layout: page
tags: [api]
title: Fyne API "widget.ImageSegment"
---

# widget.ImageSegment
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ImageSegment

```go
type ImageSegment struct {
	Source fyne.URI
	Title  string
}
```

ImageSegment represents an image within a rich text widget.


<div class="since">Since: <code>
2.3</code></div>

#### func (*ImageSegment) Inline

```go
func (i *ImageSegment) Inline() bool
```
Inline returns false as images in rich text are blocks.

#### func (*ImageSegment) Select

```go
func (i *ImageSegment) Select(begin, end fyne.Position)
```
Select tells the segment that the user is selecting the content between the two positions.

#### func (*ImageSegment) SelectedText

```go
func (i *ImageSegment) SelectedText() string
```
SelectedText should return the text representation of any content currently selected through the Select call.

#### func (*ImageSegment) Textual

```go
func (i *ImageSegment) Textual() string
```
Textual returns the content of this segment rendered to plain text.

#### func (*ImageSegment) Unselect

```go
func (i *ImageSegment) Unselect()
```
Unselect tells the segment that the user is has cancelled the previous selection.

#### func (*ImageSegment) Update

```go
func (i *ImageSegment) Update(o fyne.CanvasObject)
```
Update applies the current state of this image segment to an existing visual.

#### func (*ImageSegment) Visual

```go
func (i *ImageSegment) Visual() fyne.CanvasObject
```
Visual returns the image widget required to render this segment.
