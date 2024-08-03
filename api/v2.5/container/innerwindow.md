---
layout: page
tags: [api]
title: Fyne API "container.InnerWindow"
package: fyne.io/fyne/v2/container
---

# container.InnerWindow
---
```go
import "fyne.io/fyne/v2/container"
```

## Usage

#### type InnerWindow

```go
type InnerWindow struct {
	widget.BaseWidget

	CloseIntercept                                      func()                `json:"-"`
	OnDragged, OnResized                                func(*fyne.DragEvent) `json:"-"`
	OnMinimized, OnMaximized, OnTappedBar, OnTappedIcon func()                `json:"-"`
	Icon                                                fyne.Resource
}
```

InnerWindow defines a container that wraps content in a window border - that can then be placed inside a regular container/canvas.


<div class="since">Since: <code>
2.5</code></div>

#### func  NewInnerWindow

```go
func NewInnerWindow(title string, content fyne.CanvasObject) *InnerWindow
```
NewInnerWindow creates a new window border around the given `content`, displaying the `title` along the top. This will behave like a normal contain and will probably want to be added to a `MultipleWindows` parent.


<div class="since">Since: <code>
2.5</code></div>

#### func (*InnerWindow) Close

```go
func (w *InnerWindow) Close()
```

#### func (*InnerWindow) CreateRenderer

```go
func (w *InnerWindow) CreateRenderer() fyne.WidgetRenderer
```

#### func (*InnerWindow) SetContent

```go
func (w *InnerWindow) SetContent(obj fyne.CanvasObject)
```

#### func (*InnerWindow) SetPadded

```go
func (w *InnerWindow) SetPadded(pad bool)
```

#### func (*InnerWindow) SetTitle

```go
func (w *InnerWindow) SetTitle(title string)
```
