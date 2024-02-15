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

#### type Button

```go
type Button struct {
	DisableableWidget
	Text  string
	Style ButtonStyle
	Icon  fyne.Resource

	OnTapped func() `json:"-"`

	HideShadow bool
}
```

Button widget has a text label and triggers an event func when clicked

#### func  NewButton

```go
func NewButton(label string, tapped func()) *Button
```
NewButton creates a new button widget with the set label and tap handler

#### func  NewButtonWithIcon

```go
func NewButtonWithIcon(label string, icon fyne.Resource, tapped func()) *Button
```
NewButtonWithIcon creates a new button widget with the specified label, themed icon and tap handler

#### func (*Button) CreateRenderer

```go
func (b *Button) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Button) MinSize

```go
func (b *Button) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Button) MouseIn

```go
func (b *Button) MouseIn(*desktop.MouseEvent)
```
MouseIn is called when a desktop pointer enters the widget

#### func (*Button) MouseMoved

```go
func (b *Button) MouseMoved(*desktop.MouseEvent)
```
MouseMoved is called when a desktop pointer hovers over the widget

#### func (*Button) MouseOut

```go
func (b *Button) MouseOut()
```
MouseOut is called when a desktop pointer exits the widget

#### func (*Button) SetIcon

```go
func (b *Button) SetIcon(icon fyne.Resource)
```
SetIcon updates the icon on a label - pass nil to hide an icon

#### func (*Button) SetText

```go
func (b *Button) SetText(text string)
```
SetText allows the button label to be changed

#### func (*Button) Tapped

```go
func (b *Button) Tapped(*fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any tap handler

#### func (*Button) TappedSecondary

```go
func (b *Button) TappedSecondary(*fyne.PointEvent)
```
TappedSecondary is called when a secondary pointer tapped event is captured
