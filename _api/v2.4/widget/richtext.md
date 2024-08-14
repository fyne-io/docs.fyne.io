---
layout: page
tags: [api]
title: Fyne API "widget.RichText"
package: fyne.io/fyne/v2/widget
---

# widget.RichText
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type RichText

```go
type RichText struct {
	BaseWidget
	Segments   []RichTextSegment
	Wrapping   fyne.TextWrap
	Scroll     widget.ScrollDirection
	Truncation fyne.TextTruncation
}
```

RichText represents the base element for a rich text-based widget.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewRichText

```go
func NewRichText(segments ...RichTextSegment) *RichText
```
NewRichText returns a new RichText widget that renders the given text and segments. If no segments are specified it will be converted to a single segment using the default text settings.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewRichTextFromMarkdown

```go
func NewRichTextFromMarkdown(content string) *RichText
```
NewRichTextFromMarkdown configures a RichText widget by parsing the provided markdown content.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewRichTextWithText

```go
func NewRichTextWithText(text string) *RichText
```
NewRichTextWithText returns a new RichText widget that renders the given text. The string will be converted to a single text segment using the default text settings.


<div class="since">Since: <code>
2.1</code></div>

#### func (*RichText) CreateRenderer

```go
func (t *RichText) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*RichText) MinSize

```go
func (t *RichText) MinSize() fyne.Size
```
MinSize calculates the minimum size of a rich text widget. This is based on the contained text with a standard amount of padding added.

#### func (*RichText) ParseMarkdown

```go
func (t *RichText) ParseMarkdown(content string)
```
ParseMarkdown allows setting the content of this RichText widget from a markdown string. It will replace the content of this widget similarly to SetText, but with the appropriate formatting.

#### func (*RichText) Refresh

```go
func (t *RichText) Refresh()
```
Refresh triggers a redraw of the rich text.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*RichText) Resize

```go
func (t *RichText) Resize(size fyne.Size)
```
Resize sets a new size for the rich text. This should only be called if it is not in a container with a layout manager.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*RichText) String

```go
func (t *RichText) String() string
```
String returns the text widget buffer as string
