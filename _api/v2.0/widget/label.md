---
layout: page
tags: [api]
title: Fyne API "widget.Label"
---

# widget.Label
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
	Alignment fyne.TextAlign // The alignment of the Text
	Wrapping  fyne.TextWrap  // The wrapping of the Text
	TextStyle fyne.TextStyle // The style of the label text
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
NewLabelWithData returns an Label widget connected to the specified data source.


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
MinSize returns the size that this widget should not shrink below

#### func (*Label) Refresh

```go
func (l *Label) Refresh()
```
Refresh checks if the text content should be updated then refreshes the graphical context

#### func (*Label) Resize

```go
func (l *Label) Resize(size fyne.Size)
```
Resize sets a new size for the label. Note this should not be used if the widget is being managed by a Layout within a Container.

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
