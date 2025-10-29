---
tags: [api]
title: widget.Label
slug: label

aliases:
- /api/widget/label
- /api/widget/label.html
- /api/v2.0/widget/label
- /api/v2.0/widget/label.html
- /api/v2.1/widget/label
- /api/v2.1/widget/label.html
- /api/v2.2/widget/label
- /api/v2.2/widget/label.html
- /api/v2.3/widget/label
- /api/v2.3/widget/label.html
- /api/v2.4/widget/label
- /api/v2.4/widget/label.html
- /api/v2.5/widget/label
- /api/v2.5/widget/label.html
- /api/v2.6/widget/label
- /api/v2.6/widget/label.html
- /api/v2.7/widget/label
- /api/v2.7/widget/label.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Label

```go
type Label struct {
	BaseWidget
	Text      string
	Alignment fyne.TextAlign // The alignment of the text
	Wrapping  fyne.TextWrap  // The wrapping of the text
	TextStyle fyne.TextStyle // The style of the label text

	// The truncation mode of the text
	//
	// Since: 2.4
	Truncation fyne.TextTruncation
	// Importance informs how the label should be styled, i.e. warning or disabled
	//
	// Since: 2.4
	Importance Importance

	// The theme size name for the text size of the label
	//
	// Since: 2.6
	SizeName fyne.ThemeSizeName

	// If set to true, Selectable indicates that this label should support select interaction
	// to allow the text to be copied.
	//
	//Since: 2.6
	Selectable bool
}
```

Label widget is a label component with appropriate padding and layout.

#### func  NewLabel

```go
func NewLabel(text string) *Label
```
NewLabel creates a new label widget with the set text content

#### func  NewLabelWithData

```go
func NewLabelWithData(data binding.String) *Label
```
NewLabelWithData returns a Label widget connected to the specified data source.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewLabelWithStyle

```go
func NewLabelWithStyle(text string, alignment fyne.TextAlign, style fyne.TextStyle) *Label
```
NewLabelWithStyle creates a new label widget with the set text content

#### func (*Label) Bind

```go
func (l *Label) Bind(data binding.String)
```
Bind connects the specified data source to this Label. The current value will be displayed and any changes in the data will cause the widget to update.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Label) CreateRenderer

```go
func (l *Label) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Label) MinSize

```go
func (l *Label) MinSize() fyne.Size
```
MinSize returns the size that this label should not shrink below.

#### func (*Label) Refresh

```go
func (l *Label) Refresh()
```
Refresh triggers a redraw of the label.

#### func (*Label) SelectedText

```go
func (l *Label) SelectedText() string
```
SelectedText returns the text currently selected in this Label. If the label is not Selectable it will return an empty string. If there is no selection it will return the empty string.


<div class="since">Since: <code>
2.6</code></div>

#### func (*Label) SetText

```go
func (l *Label) SetText(text string)
```
SetText sets the text of the label

#### func (*Label) Unbind

```go
func (l *Label) Unbind()
```
Unbind disconnects any configured data source from this Label. The current value will remain at the last value of the data source.


<div class="since">Since: <code>
2.0</code></div>
