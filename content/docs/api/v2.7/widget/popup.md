---
tags: [api]
title: widget.PopUp
slug: popup

aliases:
- /api/v2.7/widget/popup

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type PopUp struct {
	BaseWidget

	Content fyne.CanvasObject
	Canvas  fyne.Canvas
}
```

PopUp is a widget that can float above the user interface. It wraps any standard elements with padding and a shadow. If it is modal then the shadow will cover the entire canvas it hovers over and block interactions.

###

```go
func NewModalPopUp(content fyne.CanvasObject, canvas fyne.Canvas) *PopUp
```
NewModalPopUp creates a new popUp for the specified content and displays it on the passed canvas. A modal PopUp blocks interactions with underlying elements, covered with a semi-transparent overlay.

###

```go
func NewPopUp(content fyne.CanvasObject, canvas fyne.Canvas) *PopUp
```
NewPopUp creates a new popUp for the specified content and displays it on the passed canvas.

###

```go
func (p *PopUp) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (p *PopUp) Hide()
```
Hide this widget, if it was previously visible

###

```go
func (p *PopUp) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

###

```go
func (p *PopUp) Move(pos fyne.Position)
```
Move the widget to a new position. A PopUp position is absolute to the top, left of its canvas. For PopUp this actually moves the content so checking Position() will not return the same value as is set here.

###

```go
func (p *PopUp) Resize(size fyne.Size)
```
Resize changes the size of the PopUp's content. PopUps always have the size of their canvas, but this call updates the size of the content portion.

###

```go
func (p *PopUp) Show()
```
Show this pop-up as overlay if not already shown.

###

```go
func (p *PopUp) ShowAtPosition(pos fyne.Position)
```
ShowAtPosition shows this pop-up at the given position.

###

```go
func (p *PopUp) ShowAtRelativePosition(rel fyne.Position, to fyne.CanvasObject)
```
ShowAtRelativePosition shows this pop-up at the given position relative to stated object.

Since 2.4

###

```go
func (p *PopUp) Tapped(e *fyne.PointEvent)
```
Tapped is called when the user taps the popUp. If not modal and the tap is outside the content area, then dismiss this widget

###

```go
func (p *PopUp) TappedSecondary(e *fyne.PointEvent)
```
TappedSecondary is called when the user right/alt taps the popUp. If not modal and the tap is outside the content area, then dismiss this widget
