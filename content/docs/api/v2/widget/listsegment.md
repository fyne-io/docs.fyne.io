---
tags: [api]
title: widget.ListSegment
slug: listsegment

aliases:
- /api/widget/listsegment
- /api/widget/listsegment.html
- /api/v2.0/widget/listsegment
- /api/v2.0/widget/listsegment.html
- /api/v2.1/widget/listsegment
- /api/v2.1/widget/listsegment.html
- /api/v2.2/widget/listsegment
- /api/v2.2/widget/listsegment.html
- /api/v2.3/widget/listsegment
- /api/v2.3/widget/listsegment.html
- /api/v2.4/widget/listsegment
- /api/v2.4/widget/listsegment.html
- /api/v2.5/widget/listsegment
- /api/v2.5/widget/listsegment.html
- /api/v2.6/widget/listsegment
- /api/v2.6/widget/listsegment.html
- /api/v2.7/widget/listsegment
- /api/v2.7/widget/listsegment.html

package: fyne.io/fyne/v2/widget
---


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

#### func (*ListSegment) SetStartNumber

```go
func (l *ListSegment) SetStartNumber(s int)
```
SetStartNumber sets the starting number for an ordered list. Unordered lists are not affected.


<div class="since">Since: <code>
2.7</code></div>

#### func (*ListSegment) StartNumber

```go
func (l *ListSegment) StartNumber() int
```
StartNumber return the starting number for an ordered list.


<div class="since">Since: <code>
2.7</code></div>

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
Update doesn't need to change a list visual.

#### func (*ListSegment) Visual

```go
func (l *ListSegment) Visual() fyne.CanvasObject
```
Visual returns no additional elements for this segment.
