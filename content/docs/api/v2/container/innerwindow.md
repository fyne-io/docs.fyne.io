---
tags: [api]
title: container.InnerWindow
slug: innerwindow

aliases:
- /api/container/innerwindow
- /api/container/innerwindow.html
- /api/v2.0/container/innerwindow
- /api/v2.0/container/innerwindow.html
- /api/v2.1/container/innerwindow
- /api/v2.1/container/innerwindow.html
- /api/v2.2/container/innerwindow
- /api/v2.2/container/innerwindow.html
- /api/v2.3/container/innerwindow
- /api/v2.3/container/innerwindow.html
- /api/v2.4/container/innerwindow
- /api/v2.4/container/innerwindow.html
- /api/v2.5/container/innerwindow
- /api/v2.5/container/innerwindow.html
- /api/v2.6/container/innerwindow
- /api/v2.6/container/innerwindow.html
- /api/v2.7/container/innerwindow
- /api/v2.7/container/innerwindow.html

package: fyne.io/fyne/v2/container
---


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

	// Alignment allows an inner window to specify if the buttons should be on the left
	// (`ButtonAlignLeading`) or right of the window border.
	//
	// Since: 2.6
	Alignment widget.ButtonAlign
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

#### func (*InnerWindow) SetMaximized

```go
func (w *InnerWindow) SetMaximized(max bool)
```
SetMaximized tells the window if the maximized state should be set or not.


<div class="since">Since: <code>
2.6</code></div>

#### func (*InnerWindow) SetPadded

```go
func (w *InnerWindow) SetPadded(pad bool)
```

#### func (*InnerWindow) SetTitle

```go
func (w *InnerWindow) SetTitle(title string)
```
