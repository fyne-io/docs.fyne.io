---
layout: page
tags: [api]
title: Fyne API "widget.Check"
package: fyne.io/fyne/v2/widget
---

# widget.Check
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Check

```go
type Check struct {
	DisableableWidget
	Text    string
	Checked bool

	// Partial check is when there is an indeterminate state (usually meaning that child items are some-what checked).
	// Turning this on will override the checked state and show a dash icon (neither checked nor unchecked).
	// The user interaction cannot turn this on, tapping a partial check state will set `Checked` to true.
	//
	// Since: 2.6
	Partial bool

	OnChanged func(bool) `json:"-"`
}
```

Check widget has a text label and a checked (or unchecked) icon and triggers an event func when toggled

#### func  NewCheck

```go
func NewCheck(label string, changed func(bool)) *Check
```
NewCheck creates a new check widget with the set label and change handler

#### func  NewCheckWithData

```go
func NewCheckWithData(label string, data binding.Bool) *Check
```
NewCheckWithData returns a check widget connected with the specified data source.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Check) Bind

```go
func (c *Check) Bind(data binding.Bool)
```
Bind connects the specified data source to this Check. The current value will be displayed and any changes in the data will cause the widget to update. User interactions with this Check will set the value into the data source.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Check) CreateRenderer

```go
func (c *Check) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Check) FocusGained

```go
func (c *Check) FocusGained()
```
FocusGained is called when the Check has been given focus.

#### func (*Check) FocusLost

```go
func (c *Check) FocusLost()
```
FocusLost is called when the Check has had focus removed.

#### func (*Check) Hide

```go
func (c *Check) Hide()
```
Hide this widget, if it was previously visible

#### func (*Check) MinSize

```go
func (c *Check) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Check) MouseIn

```go
func (c *Check) MouseIn(me *desktop.MouseEvent)
```
MouseIn is called when a desktop pointer enters the widget

#### func (*Check) MouseMoved

```go
func (c *Check) MouseMoved(me *desktop.MouseEvent)
```
MouseMoved is called when a desktop pointer hovers over the widget

#### func (*Check) MouseOut

```go
func (c *Check) MouseOut()
```
MouseOut is called when a desktop pointer exits the widget

#### func (*Check) SetChecked

```go
func (c *Check) SetChecked(checked bool)
```
SetChecked sets the checked state and refreshes widget If the `Partial` state is set this will be turned off to respect the `checked` bool passed in here.

#### func (*Check) SetText

```go
func (c *Check) SetText(text string)
```
SetText sets the text of the Check


<div class="since">Since: <code>
2.4</code></div>

#### func (*Check) Tapped

```go
func (c *Check) Tapped(pe *fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any change handler

#### func (*Check) TypedKey

```go
func (c *Check) TypedKey(key *fyne.KeyEvent)
```
TypedKey receives key input events when the Check is focused.

#### func (*Check) TypedRune

```go
func (c *Check) TypedRune(r rune)
```
TypedRune receives text input events when the Check is focused.

#### func (*Check) Unbind

```go
func (c *Check) Unbind()
```
Unbind disconnects any configured data source from this Check. The current value will remain at the last value of the data source.


<div class="since">Since: <code>
2.0</code></div>
