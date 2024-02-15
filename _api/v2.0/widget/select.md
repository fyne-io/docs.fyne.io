---
layout: page
tags: [api]
title: Fyne API "widget.Select"
---

# widget.Select
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Select

```go
type Select struct {
	DisableableWidget

	Selected    string
	Options     []string
	PlaceHolder string
	OnChanged   func(string) `json:"-"`
}
```

Select widget has a list of options, with the current one shown, and triggers an event func when clicked

#### func  NewSelect

```go
func NewSelect(options []string, changed func(string)) *Select
```
NewSelect creates a new select widget with the set list of options and changes handler

#### func (*Select) ClearSelected

```go
func (s *Select) ClearSelected()
```
ClearSelected clears the current option of the select widget. After clearing the current option, the Select widget's PlaceHolder will be displayed.

#### func (*Select) CreateRenderer

```go
func (s *Select) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Select) FocusGained

```go
func (s *Select) FocusGained()
```
FocusGained is called after this Select has gained focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Select) FocusLost

```go
func (s *Select) FocusLost()
```
FocusLost is called after this Select has lost focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Select) Hide

```go
func (s *Select) Hide()
```
Hide hides the select.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Select) MinSize

```go
func (s *Select) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Select) MouseIn

```go
func (s *Select) MouseIn(*desktop.MouseEvent)
```
MouseIn is called when a desktop pointer enters the widget

#### func (*Select) MouseMoved

```go
func (s *Select) MouseMoved(*desktop.MouseEvent)
```
MouseMoved is called when a desktop pointer hovers over the widget

#### func (*Select) MouseOut

```go
func (s *Select) MouseOut()
```
MouseOut is called when a desktop pointer exits the widget

#### func (*Select) Move

```go
func (s *Select) Move(pos fyne.Position)
```
Move changes the relative position of the select.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Select) Resize

```go
func (s *Select) Resize(size fyne.Size)
```
Resize sets a new size for a widget. Note this should not be used if the widget is being managed by a Layout within a Container.

#### func (*Select) SelectedIndex

```go
func (s *Select) SelectedIndex() int
```
SelectedIndex returns the index value of the currently selected item in Options list. It will return -1 if there is no selection.

#### func (*Select) SetSelected

```go
func (s *Select) SetSelected(text string)
```
SetSelected sets the current option of the select widget

#### func (*Select) SetSelectedIndex

```go
func (s *Select) SetSelectedIndex(index int)
```
SetSelectedIndex will set the Selected option from the value in Options list at index position.

#### func (*Select) Tapped

```go
func (s *Select) Tapped(*fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any tap handler

#### func (*Select) TypedKey

```go
func (s *Select) TypedKey(event *fyne.KeyEvent)
```
TypedKey is called if a key event happens while this Select is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Select) TypedRune

```go
func (s *Select) TypedRune(_ rune)
```
TypedRune is called if a text event happens while this Select is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>
