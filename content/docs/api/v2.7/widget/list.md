---
tags: [api]
title: widget.List
slug: list

aliases:
- /api/v2.7/widget/list

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type List struct {
	BaseWidget

	// Length is a callback for returning the number of items in the list.
	Length func() int `json:"-"`

	// CreateItem is a callback invoked to create a new widget to render
	// a row in the list.
	CreateItem func() fyne.CanvasObject `json:"-"`

	// UpdateItem is a callback invoked to update a list row widget
	// to display a new row in the list. The UpdateItem callback should
	// only update the given item, it should not invoke APIs that would
	// change other properties of the list itself.
	UpdateItem func(id ListItemID, item fyne.CanvasObject) `json:"-"`

	// OnSelected is a callback to be notified when a given item
	// in the list has been selected.
	OnSelected func(id ListItemID) `json:"-"`

	// OnSelected is a callback to be notified when a given item
	// in the list has been unselected.
	OnUnselected func(id ListItemID) `json:"-"`

	// HideSeparators hides the separators between list rows
	//
	// Since: 2.5
	HideSeparators bool
}
```

List is a widget that pools list items for performance and lays the items out in a vertical direction inside of a scroller. By default, List requires that all items are the same size, but specific rows can have their heights set with SetItemHeight.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func NewList(length func() int, createItem func() fyne.CanvasObject, updateItem func(ListItemID, fyne.CanvasObject)) *List
```
NewList creates and returns a list widget for displaying items in a vertical layout with scrolling and caching for performance.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func NewListWithData(data binding.DataList, createItem func() fyne.CanvasObject, updateItem func(binding.DataItem, fyne.CanvasObject)) *List
```
NewListWithData creates a new list widget that will display the contents of the provided data.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func (l *List) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer.

###

```go
func (l *List) FocusGained()
```
FocusGained is called after this List has gained focus.

###

```go
func (l *List) FocusLost()
```
FocusLost is called after this List has lost focus.

###

```go
func (l *List) GetScrollOffset() float32
```
GetScrollOffset returns the current scroll offset position


<div class="since">Since: <code>
2.5</code></div>

###

```go
func (l *List) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.

###

```go
func (l *List) RefreshItem(id ListItemID)
```
RefreshItem refreshes a single item, specified by the item ID passed in.


<div class="since">Since: <code>
2.4</code></div>

###

```go
func (l *List) Resize(s fyne.Size)
```
Resize is called when this list should change size. We refresh to ensure invisible items are drawn.

###

```go
func (l *List) ScrollTo(id ListItemID)
```
ScrollTo scrolls to the item represented by id


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (l *List) ScrollToBottom()
```
ScrollToBottom scrolls to the end of the list


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (l *List) ScrollToOffset(offset float32)
```
ScrollToOffset scrolls the list to the given offset position.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func (l *List) ScrollToTop()
```
ScrollToTop scrolls to the start of the list


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (l *List) Select(id ListItemID)
```
Select add the item identified by the given ID to the selection.

###

```go
func (l *List) SetItemHeight(id ListItemID, height float32)
```
SetItemHeight supports changing the height of the specified list item. Items normally take the height of the template returned from the CreateItem callback. The height parameter uses the same units as a fyne.Size type and refers to the internal content height not including the divider size.


<div class="since">Since: <code>
2.3</code></div>

###

```go
func (l *List) TypedKey(event *fyne.KeyEvent)
```
TypedKey is called if a key event happens while this List is focused.

###

```go
func (l *List) TypedRune(_ rune)
```
TypedRune is called if a text event happens while this List is focused.

###

```go
func (l *List) Unselect(id ListItemID)
```
Unselect removes the item identified by the given ID from the selection.

###

```go
func (l *List) UnselectAll()
```
UnselectAll removes all items from the selection.


<div class="since">Since: <code>
2.1</code></div>
