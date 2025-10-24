---
tags: [api]
title: widget.Select
slug: select

aliases:
- /api/v2.7/widget/select

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type Select struct {
	DisableableWidget

	// Alignment sets the text alignment of the select and its list of options.
	//
	// Since: 2.1
	Alignment   fyne.TextAlign
	Selected    string
	Options     []string
	PlaceHolder string
	OnChanged   func(string) `json:"-"`
}
```

Select widget has a list of options, with the current one shown, and triggers an event func when clicked

###

```go
func NewSelect(options []string, changed func(string)) *Select
```
NewSelect creates a new select widget with the set list of options and changes handler

###

```go
func NewSelectWithData(options []string, data binding.String) *Select
```
NewSelectWithData returns a new select widget connected to the specified data source.


<div class="since">Since: <code>
2.6</code></div>

###

```go
func (s *Select) Bind(data binding.String)
```
Bind connects the specified data source to this select. The current value will be displayed and any changes in the data will cause the widget to update.


<div class="since">Since: <code>
2.6</code></div>

###

```go
func (s *Select) ClearSelected()
```
ClearSelected clears the current option of the select widget. After clearing the current option, the Select widget's PlaceHolder will be displayed.

###

```go
func (s *Select) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (s *Select) FocusGained()
```
FocusGained is called after this Select has gained focus.

###

```go
func (s *Select) FocusLost()
```
FocusLost is called after this Select has lost focus.

###

```go
func (s *Select) Hide()
```
Hide hides the select.

###

```go
func (s *Select) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

###

```go
func (s *Select) MouseIn(*desktop.MouseEvent)
```
MouseIn is called when a desktop pointer enters the widget

###

```go
func (s *Select) MouseMoved(*desktop.MouseEvent)
```
MouseMoved is called when a desktop pointer hovers over the widget

###

```go
func (s *Select) MouseOut()
```
MouseOut is called when a desktop pointer exits the widget

###

```go
func (s *Select) Move(pos fyne.Position)
```
Move changes the relative position of the select.

###

```go
func (s *Select) Resize(size fyne.Size)
```
Resize sets a new size for a widget. Note this should not be used if the widget is being managed by a Layout within a Container.

###

```go
func (s *Select) SelectedIndex() int
```
SelectedIndex returns the index value of the currently selected item in Options list. It will return -1 if there is no selection.

###

```go
func (s *Select) SetOptions(options []string)
```
SetOptions updates the list of options available and refreshes the widget


<div class="since">Since: <code>
2.4</code></div>

###

```go
func (s *Select) SetSelected(text string)
```
SetSelected sets the current option of the select widget

###

```go
func (s *Select) SetSelectedIndex(index int)
```
SetSelectedIndex will set the Selected option from the value in Options list at index position.

###

```go
func (s *Select) Tapped(*fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any tap handler

###

```go
func (s *Select) TypedKey(event *fyne.KeyEvent)
```
TypedKey is called if a key event happens while this Select is focused.

###

```go
func (s *Select) TypedRune(_ rune)
```
TypedRune is called if a text event happens while this Select is focused.

###

```go
func (s *Select) Unbind()
```
Unbind disconnects any configured data source from this Select. The current value will remain at the last value of the data source.


<div class="since">Since: <code>
2.6</code></div>
