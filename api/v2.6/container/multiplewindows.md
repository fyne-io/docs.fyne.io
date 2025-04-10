---
layout: page
tags: [api]
title: Fyne API "container.MultipleWindows"
package: fyne.io/fyne/v2/container
---

# container.MultipleWindows
---
```go
import "fyne.io/fyne/v2/container"
```

## Usage

#### type MultipleWindows

```go
type MultipleWindows struct {
	widget.BaseWidget

	Windows []*InnerWindow
}
```

MultipleWindows is a container that handles multiple `InnerWindow` containers. Each inner window can be dragged, resized and the stacking will change when the title bar is tapped.


<div class="since">Since: <code>
2.5</code></div>

#### func  NewMultipleWindows

```go
func NewMultipleWindows(wins ...*InnerWindow) *MultipleWindows
```
NewMultipleWindows creates a new `MultipleWindows` container to manage many inner windows. The initial window list is passed optionally to this constructor function. You can add new more windows to this container by calling `Add` or updating the `Windows` field and calling `Refresh`.


<div class="since">Since: <code>
2.5</code></div>

#### func (*MultipleWindows) Add

```go
func (m *MultipleWindows) Add(w *InnerWindow)
```

#### func (*MultipleWindows) CreateRenderer

```go
func (m *MultipleWindows) CreateRenderer() fyne.WidgetRenderer
```

#### func (*MultipleWindows) Refresh

```go
func (m *MultipleWindows) Refresh()
```
