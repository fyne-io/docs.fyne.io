---
tags: [api]
title: container.Clip
slug: clip

aliases:
- /api/v2.7/container/clip

package: fyne.io/fyne/v2/container
---


---
```go
import "fyne.io/fyne/v2/container"
```

#

###

```go
type Clip struct {
	widget.BaseWidget
	Content fyne.CanvasObject
}
```

Clip describes a rectangular region that will clip anything outside its bounds.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func NewClip(content fyne.CanvasObject) *Clip
```
NewClip returns a new rectangular clipping object.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func (c *Clip) CreateRenderer() fyne.WidgetRenderer
```

###

```go
func (c *Clip) MinSize() fyne.Size
```
MinSize for a Clip simply returns Size{1, 1} as there is no explicit content
