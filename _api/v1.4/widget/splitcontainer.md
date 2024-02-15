---
layout: page
tags: [api]
title: Fyne API "widget.SplitContainer"
---

# widget.SplitContainer
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type SplitContainer

```go
type SplitContainer struct {
	BaseWidget
	Offset     float64
	Horizontal bool
	Leading    fyne.CanvasObject
	Trailing   fyne.CanvasObject
}
```

SplitContainer defines a container whose size is split between two children.


<div class="deprecated">
Deprecated: use container.Split instead.</div>

#### func  NewHSplitContainer

```go
func NewHSplitContainer(leading, trailing fyne.CanvasObject) *SplitContainer
```
NewHSplitContainer creates a horizontally arranged container with the specified leading and trailing elements. A vertical split bar that can be dragged will be added between the elements.


<div class="deprecated">
Deprecated: use container.NewHSplit instead.</div>

#### func  NewVSplitContainer

```go
func NewVSplitContainer(top, bottom fyne.CanvasObject) *SplitContainer
```
NewVSplitContainer creates a vertically arranged container with the specified top and bottom elements. A horizontal split bar that can be dragged will be added between the elements.


<div class="deprecated">
Deprecated: use container.NewVSplit instead.</div>

#### func (*SplitContainer) CreateRenderer

```go
func (s *SplitContainer) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*SplitContainer) SetOffset

```go
func (s *SplitContainer) SetOffset(offset float64)
```
SetOffset sets the offset (0.0 to 1.0) of the SplitContainer divider. 0.0 - Leading is min size, Trailing uses all remaining space. 0.5 - Leading & Trailing equally share the available space. 1.0 - Trailing is min size, Leading uses all remaining space.
