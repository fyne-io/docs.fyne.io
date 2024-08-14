---
layout: page
tags: [api]
title: Fyne API "widget.GridWrap"
package: fyne.io/fyne/v2/widget
---

# widget.GridWrap
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type GridWrap

```go
type GridWrap struct {
	BaseWidget

	Length       func() int                                      `json:"-"`
	CreateItem   func() fyne.CanvasObject                        `json:"-"`
	UpdateItem   func(id GridWrapItemID, item fyne.CanvasObject) `json:"-"`
	OnSelected   func(id GridWrapItemID)                         `json:"-"`
	OnUnselected func(id GridWrapItemID)                         `json:"-"`
}
```

GridWrap is a widget with an API very similar to widget.List, that lays out items in a scrollable wrapping grid similar to container.NewGridWrap. It caches and reuses widgets for performance.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewGridWrap

```go
func NewGridWrap(length func() int, createItem func() fyne.CanvasObject, updateItem func(GridWrapItemID, fyne.CanvasObject)) *GridWrap
```
NewGridWrap creates and returns a GridWrap widget for displaying items in a wrapping grid layout with scrolling and caching for performance.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewGridWrapWithData

```go
func NewGridWrapWithData(data binding.DataList, createItem func() fyne.CanvasObject, updateItem func(binding.DataItem, fyne.CanvasObject)) *GridWrap
```
NewGridWrapWithData creates a new GridWrap widget that will display the contents of the provided data.


<div class="since">Since: <code>
2.4</code></div>

#### func (*GridWrap) CreateRenderer

```go
func (l *GridWrap) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer.

#### func (*GridWrap) FocusGained

```go
func (l *GridWrap) FocusGained()
```
FocusGained is called after this GridWrap has gained focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*GridWrap) FocusLost

```go
func (l *GridWrap) FocusLost()
```
FocusLost is called after this GridWrap has lost focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*GridWrap) GetScrollOffset

```go
func (l *GridWrap) GetScrollOffset() float32
```
GetScrollOffset returns the current scroll offset position

#### func (*GridWrap) MinSize

```go
func (l *GridWrap) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.

#### func (*GridWrap) RefreshItem

```go
func (l *GridWrap) RefreshItem(id GridWrapItemID)
```
RefreshItem refreshes a single item, specified by the item ID passed in.


<div class="since">Since: <code>
2.4</code></div>

#### func (*GridWrap) Resize

```go
func (l *GridWrap) Resize(s fyne.Size)
```
Resize is called when this GridWrap should change size. We refresh to ensure invisible items are drawn.

#### func (*GridWrap) ScrollTo

```go
func (l *GridWrap) ScrollTo(id GridWrapItemID)
```
ScrollTo scrolls to the item represented by id

#### func (*GridWrap) ScrollToBottom

```go
func (l *GridWrap) ScrollToBottom()
```
ScrollToBottom scrolls to the end of the list

#### func (*GridWrap) ScrollToOffset

```go
func (l *GridWrap) ScrollToOffset(offset float32)
```
ScrollToOffset scrolls the list to the given offset position

#### func (*GridWrap) ScrollToTop

```go
func (l *GridWrap) ScrollToTop()
```
ScrollToTop scrolls to the start of the list

#### func (*GridWrap) Select

```go
func (l *GridWrap) Select(id GridWrapItemID)
```
Select adds the item identified by the given ID to the selection.

#### func (*GridWrap) TypedKey

```go
func (l *GridWrap) TypedKey(event *fyne.KeyEvent)
```
TypedKey is called if a key event happens while this GridWrap is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*GridWrap) TypedRune

```go
func (l *GridWrap) TypedRune(_ rune)
```
TypedRune is called if a text event happens while this GridWrap is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*GridWrap) Unselect

```go
func (l *GridWrap) Unselect(id GridWrapItemID)
```
Unselect removes the item identified by the given ID from the selection.

#### func (*GridWrap) UnselectAll

```go
func (l *GridWrap) UnselectAll()
```
UnselectAll removes all items from the selection.
