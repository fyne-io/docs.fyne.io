---
layout: page
tags: [api]
title: Fyne API "widget.ListSegment"
package: fyne.io/fyne/v2/widget
---

# widget.ListSegment
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ListSegment

```go
type ListSegment struct {
	Items   []RichTextSegment
	Ordered bool
}
```

ListSegment includes an itemised list with the content set using the Items field.


<div class="since">Since: <code>
2.1</code></div>

#### func (*ListSegment) Inline

```go
func (l *ListSegment) Inline() bool
```
Inline returns false as a list should be in a block.

#### func (*ListSegment) Segments

```go
func (l *ListSegment) Segments() []RichTextSegment
```
Segments returns the segments required to draw bullets before each item

#### func (*ListSegment) Select

```go
func (l *ListSegment) Select(_, _ fyne.Position)
```
Select does nothing for a list container.

#### func (*ListSegment) SelectedText

```go
func (l *ListSegment) SelectedText() string
```
SelectedText returns the empty string for this list.

#### func (*ListSegment) Textual

```go
func (l *ListSegment) Textual() string
```
Textual returns no content for a list as the content is in sub-segments.

#### func (*ListSegment) Unselect

```go
func (l *ListSegment) Unselect()
```
Unselect does nothing for a list container.

#### func (*ListSegment) Update

```go
func (l *ListSegment) Update(fyne.CanvasObject)
```
Update doesnt need to change a list visual.

#### func (*ListSegment) Visual

```go
func (l *ListSegment) Visual() fyne.CanvasObject
```
Visual returns no additional elements for this segment.
