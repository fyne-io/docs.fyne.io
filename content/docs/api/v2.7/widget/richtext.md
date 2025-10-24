---
tags: [api]
title: widget.RichText
slug: richtext

aliases:
- /api/v2.7/widget/richtext

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type RichText struct {
	BaseWidget
	Segments []RichTextSegment
	Wrapping fyne.TextWrap
	Scroll   fyne.ScrollDirection

	// The truncation mode of the text
	//
	// Since: 2.4
	Truncation fyne.TextTruncation
}
```

RichText represents the base element for a rich text-based widget.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func NewRichText(segments ...RichTextSegment) *RichText
```
NewRichText returns a new RichText widget that renders the given text and segments. If no segments are specified it will be converted to a single segment using the default text settings.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func NewRichTextFromMarkdown(content string) *RichText
```
NewRichTextFromMarkdown configures a RichText widget by parsing the provided markdown content.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func NewRichTextWithText(text string) *RichText
```
NewRichTextWithText returns a new RichText widget that renders the given text. The string will be converted to a single text segment using the default text settings.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (t *RichText) AppendMarkdown(content string)
```
AppendMarkdown parses the given markdown string and appends the content to the widget, with the appropriate formatting. This API is intended for appending complete markdown documents or standalone fragments, and should not be used to parse a single markdown document piecewise.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func (t *RichText) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (t *RichText) MinSize() fyne.Size
```
MinSize calculates the minimum size of a rich text widget. This is based on the contained text with a standard amount of padding added.

###

```go
func (t *RichText) ParseMarkdown(content string)
```
ParseMarkdown allows setting the content of this RichText widget from a markdown string. It will replace the content of this widget similarly to SetText, but with the appropriate formatting.

###

```go
func (t *RichText) Refresh()
```
Refresh triggers a redraw of the rich text.

###

```go
func (t *RichText) Resize(size fyne.Size)
```
Resize sets a new size for the rich text. This should only be called if it is not in a container with a layout manager.

###

```go
func (t *RichText) String() string
```
String returns the text widget buffer as string
