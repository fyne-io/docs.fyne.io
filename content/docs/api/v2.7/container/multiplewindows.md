---
tags: [api]
title: container.MultipleWindows
slug: multiplewindows

aliases:
- /api/v2.7/container/multiplewindows

package: fyne.io/fyne/v2/container
---


---
```go
import "fyne.io/fyne/v2/container"
```

#

###

```go
type MultipleWindows struct {
	widget.BaseWidget

	Windows []*InnerWindow
}
```

MultipleWindows is a container that handles multiple `InnerWindow` containers. Each inner window can be dragged, resized and the stacking will change when the title bar is tapped.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func NewMultipleWindows(wins ...*InnerWindow) *MultipleWindows
```
NewMultipleWindows creates a new `MultipleWindows` container to manage many inner windows. The initial window list is passed optionally to this constructor function. You can add new more windows to this container by calling `Add` or updating the `Windows` field and calling `Refresh`.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func (m *MultipleWindows) Add(w *InnerWindow)
```

###

```go
func (m *MultipleWindows) CreateRenderer() fyne.WidgetRenderer
```

###

```go
func (m *MultipleWindows) Refresh()
```
