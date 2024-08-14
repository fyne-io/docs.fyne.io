---
layout: page
tags: [api]
title: Fyne API "widget.Hyperlink"
package: fyne.io/fyne/v2/widget
---

# widget.Hyperlink
---
```go
import "fyne.io/fyne/v2/widget"
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

	// OnTapped overrides the default `fyne.OpenURL` call when the link is tapped
	//
	// Since: 2.2
	OnTapped func() `json:"-"`
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

#### func (*Hyperlink) FocusGained

```go
func (hl *Hyperlink) FocusGained()
```
FocusGained is a hook called by the focus handling logic after this object gained the focus.

#### func (*Hyperlink) FocusLost

```go
func (hl *Hyperlink) FocusLost()
```
FocusLost is a hook called by the focus handling logic after this object lost the focus.

#### func (*Hyperlink) MinSize

```go
func (hl *Hyperlink) MinSize() fyne.Size
```
MinSize returns the smallest size this widget can shrink to

#### func (*Hyperlink) MouseIn

```go
func (hl *Hyperlink) MouseIn(*desktop.MouseEvent)
```
MouseIn is a hook that is called if the mouse pointer enters the element.

#### func (*Hyperlink) MouseMoved

```go
func (hl *Hyperlink) MouseMoved(*desktop.MouseEvent)
```
MouseMoved is a hook that is called if the mouse pointer moved over the element.

#### func (*Hyperlink) MouseOut

```go
func (hl *Hyperlink) MouseOut()
```
MouseOut is a hook that is called if the mouse pointer leaves the element.

#### func (*Hyperlink) Refresh

```go
func (hl *Hyperlink) Refresh()
```
Refresh triggers a redraw of the hyperlink.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

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

#### func (*Hyperlink) TypedKey

```go
func (hl *Hyperlink) TypedKey(ev *fyne.KeyEvent)
```
TypedKey is a hook called by the input handling logic on key events if this object is focused.

#### func (*Hyperlink) TypedRune

```go
func (hl *Hyperlink) TypedRune(rune)
```
TypedRune is a hook called by the input handling logic on text input events if this object is focused.
