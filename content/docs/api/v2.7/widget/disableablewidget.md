---
tags: [api]
title: widget.DisableableWidget
slug: disableablewidget

aliases:
- /api/v2.7/widget/disableablewidget

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type DisableableWidget struct {
	BaseWidget
}
```

DisableableWidget describes an extension to BaseWidget which can be disabled. Disabled widgets should have a visually distinct style when disabled, normally using theme.DisabledButtonColor.

###

```go
func (w *DisableableWidget) Disable()
```
Disable this widget so that it cannot be interacted with, updating any style appropriately.

###

```go
func (w *DisableableWidget) Disabled() bool
```
Disabled returns true if this widget is currently disabled or false if it can currently be interacted with.

###

```go
func (w *DisableableWidget) Enable()
```
Enable this widget, updating any style or features appropriately.
