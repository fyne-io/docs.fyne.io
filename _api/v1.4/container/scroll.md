---
layout: page
tags: [api]
title: Fyne API "container.Scroll"
---

# container.Scroll
---
```go
import "fyne.io/fyne/container"
```

## Usage

#### type Scroll

```go
type Scroll = widget.ScrollContainer
```

Scroll defines a container that is smaller than the Content. The Offset is used to determine the position of the child widgets within the container.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewHScroll

```go
func NewHScroll(content fyne.CanvasObject) *Scroll
```
NewHScroll creates a scrollable parent wrapping the specified content. Note that this may cause the MinSize.Width to be smaller than that of the passed object.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewScroll

```go
func NewScroll(content fyne.CanvasObject) *Scroll
```
NewScroll creates a scrollable parent wrapping the specified content. Note that this may cause the MinSize to be smaller than that of the passed object.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewVScroll

```go
func NewVScroll(content fyne.CanvasObject) *Scroll
```
NewVScroll creates a scrollable parent wrapping the specified content. Note that this may cause the MinSize.Height to be smaller than that of the passed object.


<div class="since">Since: <code>
1.4</code></div>
