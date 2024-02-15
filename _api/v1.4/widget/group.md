---
layout: page
tags: [api]
title: Fyne API "widget.Group"
---

# widget.Group
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type Group

```go
type Group struct {
	BaseWidget

	Text string
}
```

Group widget contains a list of widgets that are grouped under a dividing line and title at the top.

#### func  NewGroup

```go
func NewGroup(title string, children ...fyne.CanvasObject) *Group
```
NewGroup creates a new grouped list widget with a title and the specified list of child objects.

<div class="deprecated"> Deprecated: Consider using the Card instead.</div>

#### func  NewGroupWithScroller

```go
func NewGroupWithScroller(title string, children ...fyne.CanvasObject) *Group
```
NewGroupWithScroller creates a new grouped list widget with a title and the specified list of child objects. This group will scroll when the available space is less than needed to display the items it contains.

<div class="deprecated"> Deprecated: Consider using the Card instead.</div>

#### func (*Group) Append

```go
func (g *Group) Append(object fyne.CanvasObject)
```
Append adds a new CanvasObject to the end of the group

#### func (*Group) CreateRenderer

```go
func (g *Group) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Group) MinSize

```go
func (g *Group) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Group) Prepend

```go
func (g *Group) Prepend(object fyne.CanvasObject)
```
Prepend inserts a new CanvasObject at the top of the group
