---
layout: page
tags: [api]
title: Fyne API "container.Split"
---

# container.Split
---
```go
import "fyne.io/fyne/container"
```

## Usage

#### type Split

```go
type Split = widget.SplitContainer
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
