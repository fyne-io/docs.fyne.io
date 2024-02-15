---
layout: page
tags: [api]
title: Fyne API "widget.List"
---

# widget.List
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type List

```go
type List struct {
	BaseWidget

	Length       func() int
	CreateItem   func() fyne.CanvasObject
	UpdateItem   func(id ListItemID, item fyne.CanvasObject)
	OnSelected   func(id ListItemID)
	OnUnselected func(id ListItemID)
}
```

List is a widget that pools list items for performance and lays the items out in a vertical direction inside of a scroller. List requires that all items are the same size.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewList

```go
func NewList(length func() int, createItem func() fyne.CanvasObject, updateItem func(ListItemID, fyne.CanvasObject)) *List
```
NewList creates and returns a list widget for displaying items in a vertical layout with scrolling and caching for performance.


<div class="since">Since: <code>
1.4</code></div>

#### func (*List) CreateRenderer

```go
func (l *List) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer.

#### func (*List) MinSize

```go
func (l *List) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.

#### func (*List) Select

```go
func (l *List) Select(id ListItemID)
```
Select add the item identified by the given ID to the selection.

#### func (*List) Unselect

```go
func (l *List) Unselect(id ListItemID)
```
Unselect removes the item identified by the given ID from the selection.
