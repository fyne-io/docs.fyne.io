---
layout: page
tags: [api]
title: Fyne API "widget.ParagraphSegment"
---

# widget.ParagraphSegment
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ParagraphSegment

```go
type ParagraphSegment struct {
	Texts []RichTextSegment
}
```

ParagraphSegment wraps a number of text elements in a paragraph. It is similar to using a list of text elements when the final style is RichTextStyleParagraph.


<div class="since">Since: <code>
2.1</code></div>

#### func (*ParagraphSegment) Inline

```go
func (p *ParagraphSegment) Inline() bool
```
Inline returns false as a paragraph should be in a block.

#### func (*ParagraphSegment) Segments

```go
func (p *ParagraphSegment) Segments() []RichTextSegment
```
Segments returns the list of text elements in this paragraph.

#### func (*ParagraphSegment) Select

```go
func (p *ParagraphSegment) Select(_, _ fyne.Position)
```
Select does nothing for a paragraph container.

#### func (*ParagraphSegment) SelectedText

```go
func (p *ParagraphSegment) SelectedText() string
```
SelectedText returns the empty string for this paragraph container.

#### func (*ParagraphSegment) Textual

```go
func (p *ParagraphSegment) Textual() string
```
Textual returns no content for a paragraph container.

#### func (*ParagraphSegment) Unselect

```go
func (p *ParagraphSegment) Unselect()
```
Unselect does nothing for a paragraph container.

#### func (*ParagraphSegment) Update

```go
func (p *ParagraphSegment) Update(fyne.CanvasObject)
```
Update doesnt need to change a paragraph container.

#### func (*ParagraphSegment) Visual

```go
func (p *ParagraphSegment) Visual() fyne.CanvasObject
```
Visual returns the no extra elements.
