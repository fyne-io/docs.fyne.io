---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
---
```go
import "fyne.io/fyne/widget"
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

#### func  NewLabelWithStyle

```go
func NewLabelWithStyle(text string, alignment fyne.TextAlign, style fyne.TextStyle) *Label
```
NewLabelWithStyle creates a new label widget with the set text content

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
