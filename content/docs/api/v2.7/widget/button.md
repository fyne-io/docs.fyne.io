---
tags: [api]
title: widget.Button
slug: button

aliases:
- /api/v2.7/widget/button

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type Button struct {
	DisableableWidget
	Text string
	Icon fyne.Resource
	// Specify how prominent the button should be, High will highlight the button and Low will remove some decoration.
	//
	// Since: 1.4
	Importance    Importance
	Alignment     ButtonAlign
	IconPlacement ButtonIconPlacement

	OnTapped func() `json:"-"`
}
```

Button widget has a text label and triggers an event func when clicked

###

```go
func NewButton(label string, tapped func()) *Button
```
NewButton creates a new button widget with the set label and tap handler

###

```go
func NewButtonWithIcon(label string, icon fyne.Resource, tapped func()) *Button
```
NewButtonWithIcon creates a new button widget with the specified label, themed icon and tap handler

###

```go
func (b *Button) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (b *Button) Cursor() desktop.Cursor
```
Cursor returns the cursor type of this widget

###

```go
func (b *Button) FocusGained()
```
FocusGained is a hook called by the focus handling logic after this object gained the focus.

###

```go
func (b *Button) FocusLost()
```
FocusLost is a hook called by the focus handling logic after this object lost the focus.

###

```go
func (b *Button) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

###

```go
func (b *Button) MouseIn(*desktop.MouseEvent)
```
MouseIn is called when a desktop pointer enters the widget

###

```go
func (b *Button) MouseMoved(*desktop.MouseEvent)
```
MouseMoved is called when a desktop pointer hovers over the widget

###

```go
func (b *Button) MouseOut()
```
MouseOut is called when a desktop pointer exits the widget

###

```go
func (b *Button) SetIcon(icon fyne.Resource)
```
SetIcon updates the icon on a label - pass nil to hide an icon

###

```go
func (b *Button) SetText(text string)
```
SetText allows the button label to be changed

###

```go
func (b *Button) Tapped(*fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any tap handler

###

```go
func (b *Button) TypedKey(ev *fyne.KeyEvent)
```
TypedKey is a hook called by the input handling logic on key events if this object is focused.

###

```go
func (b *Button) TypedRune(rune)
```
TypedRune is a hook called by the input handling logic on text input events if this object is focused.
