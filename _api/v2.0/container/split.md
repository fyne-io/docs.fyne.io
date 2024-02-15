---
layout: page
tags: [api]
title: Fyne API "container.Split"
---

# container.Split
---
```go
import "fyne.io/fyne/v2/container"
```

## Usage

#### type Split

```go
type Split struct {
	widget.BaseWidget
	Offset     float64
	Horizontal bool
	Leading    fyne.CanvasObject
	Trailing   fyne.CanvasObject
}
```

Split defines a container whose size is split between two children.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewHSplit

```go
func NewHSplit(leading, trailing fyne.CanvasObject) *Split
```
NewHSplit creates a horizontally arranged container with the specified leading and trailing elements. A vertical split bar that can be dragged will be added between the elements.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewVSplit

```go
func NewVSplit(top, bottom fyne.CanvasObject) *Split
```
NewVSplit creates a vertically arranged container with the specified top and bottom elements. A horizontal split bar that can be dragged will be added between the elements.


<div class="since">Since: <code>
1.4</code></div>

#### func (*Split) CreateRenderer

```go
func (s *Split) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Split) ExtendBaseWidget

```go
func (s *Split) ExtendBaseWidget(wid fyne.Widget)
```
ExtendBaseWidget is used by an extending widget to make use of BaseWidget functionality.


<div class="deprecated">
Deprecated: Support for extending containers is being removed</div>

#### func (*Split) SetOffset

```go
func (s *Split) SetOffset(offset float64)
```
SetOffset sets the offset (0.0 to 1.0) of the Split divider. 0.0 - Leading is min size, Trailing uses all remaining space. 0.5 - Leading & Trailing equally share the available space. 1.0 - Trailing is min size, Leading uses all remaining space.
