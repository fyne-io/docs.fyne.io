---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
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

#### func  NewHSplitContainer

```go
func NewHSplitContainer(left, right fyne.CanvasObject) *SplitContainer
```
NewHSplitContainer create a splitable parent wrapping the specified children.

#### func  NewVSplitContainer

```go
func NewVSplitContainer(top, bottom fyne.CanvasObject) *SplitContainer
```
NewVSplitContainer create a splitable parent wrapping the specified children.

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
