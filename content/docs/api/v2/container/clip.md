---
tags: [api]
title: container.Clip
slug: clip

aliases:
- /api/v2/container/clip.html
- /api/v2.0/container/clip
- /api/v2.0/container/clip.html
- /api/v2.1/container/clip
- /api/v2.1/container/clip.html
- /api/v2.2/container/clip
- /api/v2.2/container/clip.html
- /api/v2.3/container/clip
- /api/v2.3/container/clip.html
- /api/v2.4/container/clip
- /api/v2.4/container/clip.html
- /api/v2.5/container/clip
- /api/v2.5/container/clip.html
- /api/v2.6/container/clip
- /api/v2.6/container/clip.html
- /api/v2.7/container/clip
- /api/v2.7/container/clip.html

package: fyne.io/fyne/v2/container
---


---
```go
import "fyne.io/fyne/v2/container"
```

## Usage

#### type Clip

```go
type Clip struct {
	widget.BaseWidget
	Content fyne.CanvasObject
}
```

Clip describes a rectangular region that will clip anything outside its bounds.


<div class="since">Since: <code>
2.7</code></div>

#### func  NewClip

```go
func NewClip(content fyne.CanvasObject) *Clip
```
NewClip returns a new rectangular clipping object.


<div class="since">Since: <code>
2.7</code></div>

#### func (*Clip) CreateRenderer

```go
func (c *Clip) CreateRenderer() fyne.WidgetRenderer
```

#### func (*Clip) MinSize

```go
func (c *Clip) MinSize() fyne.Size
```
MinSize for a Clip simply returns Size{1, 1} as there is no explicit content
