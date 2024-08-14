---
layout: page
tags: [api]
title: Fyne API "widget.BaseWidget"
package: fyne.io/fyne/v2/widget
---

# widget.BaseWidget
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type BaseWidget

```go
type BaseWidget struct {
	Hidden bool
}
```

BaseWidget provides a helper that handles basic widget behaviours.

#### func (*BaseWidget) ExtendBaseWidget

```go
func (w *BaseWidget) ExtendBaseWidget(wid fyne.Widget)
```
ExtendBaseWidget is used by an extending widget to make use of BaseWidget functionality.

#### func (*BaseWidget) Hide

```go
func (w *BaseWidget) Hide()
```
Hide this widget so it is no longer visible

#### func (*BaseWidget) MinSize

```go
func (w *BaseWidget) MinSize() fyne.Size
```
MinSize for the widget - it should never be resized below this value.

#### func (*BaseWidget) Move

```go
func (w *BaseWidget) Move(pos fyne.Position)
```
Move the widget to a new position, relative to its parent. Note this should not be used if the widget is being managed by a Layout within a Container.

#### func (*BaseWidget) Position

```go
func (w *BaseWidget) Position() fyne.Position
```
Position gets the current position of this widget, relative to its parent.

#### func (*BaseWidget) Refresh

```go
func (w *BaseWidget) Refresh()
```
Refresh causes this widget to be redrawn in it's current state

#### func (*BaseWidget) Resize

```go
func (w *BaseWidget) Resize(size fyne.Size)
```
Resize sets a new size for a widget. Note this should not be used if the widget is being managed by a Layout within a Container.

#### func (*BaseWidget) Show

```go
func (w *BaseWidget) Show()
```
Show this widget so it becomes visible

#### func (*BaseWidget) Size

```go
func (w *BaseWidget) Size() fyne.Size
```
Size gets the current size of this widget.

#### func (*BaseWidget) Visible

```go
func (w *BaseWidget) Visible() bool
```
Visible returns whether or not this widget should be visible. Note that this may not mean it is currently visible if a parent has been hidden.
