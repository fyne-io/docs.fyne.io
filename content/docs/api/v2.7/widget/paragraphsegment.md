---
tags: [api]
title: widget.ParagraphSegment
slug: paragraphsegment

aliases:
- /api/v2.7/widget/paragraphsegment

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type ParagraphSegment struct {
	Texts []RichTextSegment
}
```

ParagraphSegment wraps a number of text elements in a paragraph. It is similar to using a list of text elements when the final style is RichTextStyleParagraph.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (p *ParagraphSegment) Inline() bool
```
Inline returns false as a paragraph should be in a block.

###

```go
func (p *ParagraphSegment) Segments() []RichTextSegment
```
Segments returns the list of text elements in this paragraph.

###

```go
func (p *ParagraphSegment) Select(_, _ fyne.Position)
```
Select does nothing for a paragraph container.

###

```go
func (p *ParagraphSegment) SelectedText() string
```
SelectedText returns the empty string for this paragraph container.

###

```go
func (p *ParagraphSegment) Textual() string
```
Textual returns no content for a paragraph container.

###

```go
func (p *ParagraphSegment) Unselect()
```
Unselect does nothing for a paragraph container.

###

```go
func (p *ParagraphSegment) Update(fyne.CanvasObject)
```
Update doesn't need to change a paragraph container.

###

```go
func (p *ParagraphSegment) Visual() fyne.CanvasObject
```
Visual returns the no extra elements.
