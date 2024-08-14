---
layout: page
tags: [api]
title: Fyne API "widget.List"
package: fyne.io/fyne/v2/widget
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

	Length       func() int                                  `json:"-"`
	CreateItem   func() fyne.CanvasObject                    `json:"-"`
	UpdateItem   func(id ListItemID, item fyne.CanvasObject) `json:"-"`
	OnSelected   func(id ListItemID)                         `json:"-"`
	OnUnselected func(id ListItemID)                         `json:"-"`
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

#### func (*List) FocusGained

```go
func (l *List) FocusGained()
```
FocusGained is called after this List has gained focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*List) FocusLost

```go
func (l *List) FocusLost()
```
FocusLost is called after this List has lost focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*List) MinSize

```go
func (l *List) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.

#### func (*List) RefreshItem

```go
func (l *List) RefreshItem(id ListItemID)
```
RefreshItem refreshes a single item, specified by the item ID passed in.


<div class="since">Since: <code>
2.4</code></div>

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

#### func (*List) SetItemHeight

```go
func (l *List) SetItemHeight(id ListItemID, height float32)
```
SetItemHeight supports changing the height of the specified list item. Items normally take the height of the template returned from the CreateItem callback. The height parameter uses the same units as a fyne.Size type and refers to the internal content height not including the divider size.


<div class="since">Since: <code>
2.3</code></div>

#### func (*List) TypedKey

```go
func (l *List) TypedKey(event *fyne.KeyEvent)
```
TypedKey is called if a key event happens while this List is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*List) TypedRune

```go
func (l *List) TypedRune(_ rune)
```
TypedRune is called if a text event happens while this List is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

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
