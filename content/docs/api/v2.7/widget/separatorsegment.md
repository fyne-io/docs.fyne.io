---
tags: [api]
title: widget.SeparatorSegment
slug: separatorsegment

aliases:
- /api/v2.7/widget/separatorsegment

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type SeparatorSegment struct {
}
```

SeparatorSegment includes a horizontal separator in a rich text widget.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (s *SeparatorSegment) Inline() bool
```
Inline returns false as a separator should be full width.

###

```go
func (s *SeparatorSegment) Select(_, _ fyne.Position)
```
Select does nothing for a separator.

###

```go
func (s *SeparatorSegment) SelectedText() string
```
SelectedText returns the empty string for this separator.

###

```go
func (s *SeparatorSegment) Textual() string
```
Textual returns no content for a separator element.

###

```go
func (s *SeparatorSegment) Unselect()
```
Unselect does nothing for a separator.

###

```go
func (s *SeparatorSegment) Update(fyne.CanvasObject)
```
Update doesn't need to change a separator visual.

###

```go
func (s *SeparatorSegment) Visual() fyne.CanvasObject
```
Visual returns a new instance of a separator widget for this segment.
