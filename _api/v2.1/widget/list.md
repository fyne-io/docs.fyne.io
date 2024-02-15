---
layout: page
tags: [api]
title: Fyne API "widget.List"
---

# widget.List
---
```go
import "fyne.io/fyne/v2/widget"
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

#### func  NewListWithData

```go
func NewListWithData(data binding.DataList, createItem func() fyne.CanvasObject, updateItem func(binding.DataItem, fyne.CanvasObject)) *List
```
NewListWithData creates a new list widget that will display the contents of the provided data.


<div class="since">Since: <code>
2.0</code></div>

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

#### func (*List) Resize

```go
func (l *List) Resize(s fyne.Size)
```
Resize is called when this list should change size. We refresh to ensure invisible items are drawn.

#### func (*List) ScrollTo

```go
func (l *List) ScrollTo(id ListItemID)
```
ScrollTo scrolls to the item represented by id


<div class="since">Since: <code>
2.1</code></div>

#### func (*List) ScrollToBottom

```go
func (l *List) ScrollToBottom()
```
ScrollToBottom scrolls to the end of the list


<div class="since">Since: <code>
2.1</code></div>

#### func (*List) ScrollToTop

```go
func (l *List) ScrollToTop()
```
ScrollToTop scrolls to the start of the list


<div class="since">Since: <code>
2.1</code></div>

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

#### func (*List) UnselectAll

```go
func (l *List) UnselectAll()
```
UnselectAll removes all items from the selection.


<div class="since">Since: <code>
2.1</code></div>
