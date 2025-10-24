---
tags: [api]
title: widget.ListSegment
slug: listsegment

aliases:
- /api/v2.7/widget/listsegment

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type ListSegment struct {
	Items   []RichTextSegment
	Ordered bool
}
```

ListSegment includes an itemised list with the content set using the Items field.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (l *ListSegment) Inline() bool
```
Inline returns false as a list should be in a block.

###

```go
func (l *ListSegment) Segments() []RichTextSegment
```
Segments returns the segments required to draw bullets before each item

###

```go
func (l *ListSegment) Select(_, _ fyne.Position)
```
Select does nothing for a list container.

###

```go
func (l *ListSegment) SelectedText() string
```
SelectedText returns the empty string for this list.

###

```go
func (l *ListSegment) SetStartNumber(s int)
```
SetStartNumber sets the starting number for an ordered list. Unordered lists are not affected.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func (l *ListSegment) StartNumber() int
```
StartNumber return the starting number for an ordered list.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func (l *ListSegment) Textual() string
```
Textual returns no content for a list as the content is in sub-segments.

###

```go
func (l *ListSegment) Unselect()
```
Unselect does nothing for a list container.

###

```go
func (l *ListSegment) Update(fyne.CanvasObject)
```
Update doesn't need to change a list visual.

###

```go
func (l *ListSegment) Visual() fyne.CanvasObject
```
Visual returns no additional elements for this segment.
