---
layout: page
tags: [api]
title: Fyne API "widget.Hyperlink"
---

# widget.Hyperlink
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type Hyperlink

```go
type Hyperlink struct {
	BaseWidget
	Text      string
	URL       *url.URL
	Alignment fyne.TextAlign // The alignment of the Text
	Wrapping  fyne.TextWrap  // The wrapping of the Text
	TextStyle fyne.TextStyle // The style of the hyperlink text
}
```

Hyperlink widget is a text component with appropriate padding and layout. When clicked, the default web browser should open with a URL

#### func  NewHyperlink

```go
func NewHyperlink(text string, url *url.URL) *Hyperlink
```
NewHyperlink creates a new hyperlink widget with the set text content

#### func  NewHyperlinkWithStyle

```go
func NewHyperlinkWithStyle(text string, url *url.URL, alignment fyne.TextAlign, style fyne.TextStyle) *Hyperlink
```
NewHyperlinkWithStyle creates a new hyperlink widget with the set text content

#### func (*Hyperlink) CreateRenderer

```go
func (hl *Hyperlink) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Hyperlink) Cursor

```go
func (hl *Hyperlink) Cursor() desktop.Cursor
```
Cursor returns the cursor type of this widget

#### func (*Hyperlink) MinSize

```go
func (hl *Hyperlink) MinSize() fyne.Size
```
MinSize returns the smallest size this widget can shrink to

#### func (*Hyperlink) Resize

```go
func (hl *Hyperlink) Resize(size fyne.Size)
```
Resize sets a new size for the hyperlink. Note this should not be used if the widget is being managed by a Layout within a Container.

#### func (*Hyperlink) SetText

```go
func (hl *Hyperlink) SetText(text string)
```
SetText sets the text of the hyperlink

#### func (*Hyperlink) SetURL

```go
func (hl *Hyperlink) SetURL(url *url.URL)
```
SetURL sets the URL of the hyperlink, taking in a URL type

#### func (*Hyperlink) SetURLFromString

```go
func (hl *Hyperlink) SetURLFromString(str string) error
```
SetURLFromString sets the URL of the hyperlink, taking in a string type

#### func (*Hyperlink) Tapped

```go
func (hl *Hyperlink) Tapped(*fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any change handler
