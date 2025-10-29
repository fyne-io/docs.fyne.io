---
tags: [api]
title: widget.SeparatorSegment
slug: separatorsegment

aliases:
- /api/widget/separatorsegment
- /api/widget/separatorsegment.html
- /api/v2.0/widget/separatorsegment
- /api/v2.0/widget/separatorsegment.html
- /api/v2.1/widget/separatorsegment
- /api/v2.1/widget/separatorsegment.html
- /api/v2.2/widget/separatorsegment
- /api/v2.2/widget/separatorsegment.html
- /api/v2.3/widget/separatorsegment
- /api/v2.3/widget/separatorsegment.html
- /api/v2.4/widget/separatorsegment
- /api/v2.4/widget/separatorsegment.html
- /api/v2.5/widget/separatorsegment
- /api/v2.5/widget/separatorsegment.html
- /api/v2.6/widget/separatorsegment
- /api/v2.6/widget/separatorsegment.html
- /api/v2.7/widget/separatorsegment
- /api/v2.7/widget/separatorsegment.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type SeparatorSegment

```go
type SeparatorSegment struct {
}
```

SeparatorSegment includes a horizontal separator in a rich text widget.


<div class="since">Since: <code>
2.1</code></div>

#### func (*SeparatorSegment) Inline

```go
func (s *SeparatorSegment) Inline() bool
```
Inline returns false as a separator should be full width.

#### func (*SeparatorSegment) Select

```go
func (s *SeparatorSegment) Select(_, _ fyne.Position)
```
Select does nothing for a separator.

#### func (*SeparatorSegment) SelectedText

```go
func (s *SeparatorSegment) SelectedText() string
```
SelectedText returns the empty string for this separator.

#### func (*SeparatorSegment) Textual

```go
func (s *SeparatorSegment) Textual() string
```
Textual returns no content for a separator element.

#### func (*SeparatorSegment) Unselect

```go
func (s *SeparatorSegment) Unselect()
```
Unselect does nothing for a separator.

#### func (*SeparatorSegment) Update

```go
func (s *SeparatorSegment) Update(fyne.CanvasObject)
```
Update doesn't need to change a separator visual.

#### func (*SeparatorSegment) Visual

```go
func (s *SeparatorSegment) Visual() fyne.CanvasObject
```
Visual returns a new instance of a separator widget for this segment.
