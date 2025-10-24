---
tags: [api]
title: widget.BaseWidget
slug: basewidget

aliases:
- /api/v2.7/widget/basewidget

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type BaseWidget struct {
	Hidden bool
}
```

BaseWidget provides a helper that handles basic widget behaviours.

###

```go
func (w *BaseWidget) ExtendBaseWidget(wid fyne.Widget)
```
ExtendBaseWidget is used by an extending widget to make use of BaseWidget functionality.

###

```go
func (w *BaseWidget) Hide()
```
Hide this widget so it is no longer visible

###

```go
func (w *BaseWidget) MinSize() fyne.Size
```
MinSize for the widget - it should never be resized below this value.

###

```go
func (w *BaseWidget) Move(pos fyne.Position)
```
Move the widget to a new position, relative to its parent. Note this should not be used if the widget is being managed by a Layout within a Container.

###

```go
func (w *BaseWidget) Position() fyne.Position
```
Position gets the current position of this widget, relative to its parent.

###

```go
func (w *BaseWidget) Refresh()
```
Refresh causes this widget to be redrawn in its current state

###

```go
func (w *BaseWidget) Resize(size fyne.Size)
```
Resize sets a new size for a widget. Note this should not be used if the widget is being managed by a Layout within a Container.

###

```go
func (w *BaseWidget) Show()
```
Show this widget so it becomes visible

###

```go
func (w *BaseWidget) Size() fyne.Size
```
Size gets the current size of this widget.

###

```go
func (w *BaseWidget) Theme() fyne.Theme
```
Theme returns a cached Theme instance for this widget (or its extending widget). This will be the app theme in most cases, or a widget specific theme if it is inside a ThemeOverride container.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func (w *BaseWidget) Visible() bool
```
Visible returns whether or not this widget should be visible. Note that this may not mean it is currently visible if a parent has been hidden.
