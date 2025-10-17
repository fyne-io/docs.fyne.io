---
layout: page
tags: [api]
title: Fyne API "container.Clip"
package: fyne.io/fyne/v2/container
---

# container.Clip
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
