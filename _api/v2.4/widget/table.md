---
layout: page
tags: [api]
title: Fyne API "widget.Table"
package: fyne.io/fyne/v2/widget
---

# widget.Table
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Table

```go
type Table struct {
	BaseWidget

	Length       func() (int, int)                                `json:"-"`
	CreateCell   func() fyne.CanvasObject                         `json:"-"`
	UpdateCell   func(id TableCellID, template fyne.CanvasObject) `json:"-"`
	OnSelected   func(id TableCellID)                             `json:"-"`
	OnUnselected func(id TableCellID)                             `json:"-"`

	// ShowHeaderRow specifies that a row should be added to the table with header content.
	// This will default to an A-Z style content, unless overridden with `CreateHeader` and `UpdateHeader` calls.
	//
	// Since: 2.4
	ShowHeaderRow bool

	// ShowHeaderColumn specifies that a column should be added to the table with header content.
	// This will default to an 1-10 style numeric content, unless overridden with `CreateHeader` and `UpdateHeader` calls.
	//
	// Since: 2.4
	ShowHeaderColumn bool

	// CreateHeader is an optional function that allows overriding of the default header widget.
	// Developers must also override `UpdateHeader`.
	//
	// Since: 2.4
	CreateHeader func() fyne.CanvasObject `json:"-"`

	// UpdateHeader is used with `CreateHeader` to support custom header content.
	// The `id` parameter will have `-1` value to indicate a header, and `> 0` where the column or row refer to data.
	//
	// Since: 2.4
	UpdateHeader func(id TableCellID, template fyne.CanvasObject) `json:"-"`

	// StickyRowCount specifies how many data rows should not scroll when the content moves.
	// If `ShowHeaderRow` us `true` then the stuck row will appear immediately underneath.
	//
	// Since: 2.4
	StickyRowCount int

	// StickyColumnCount specifies how many data columns should not scroll when the content moves.
	// If `ShowHeaderColumn` us `true` then the stuck column will appear immediately next to the header.
	//
	// Since: 2.4
	StickyColumnCount int
}
```

Table widget is a grid of items that can be scrolled and a cell selected. Its performance is provided by caching cell templates created with CreateCell and re-using them with UpdateCell. The size of the content rows/columns is returned by the Length callback.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewTable

```go
func NewTable(length func() (int, int), create func() fyne.CanvasObject, update func(TableCellID, fyne.CanvasObject)) *Table
```
NewTable returns a new performant table widget defined by the passed functions. The first returns the data size in rows and columns, second parameter is a function that returns cell template objects that can be cached and the third is used to apply data at specified data location to the passed template CanvasObject.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewTableWithHeaders

```go
func NewTableWithHeaders(length func() (int, int), create func() fyne.CanvasObject, update func(TableCellID, fyne.CanvasObject)) *Table
```
NewTableWithHeaders returns a new performant table widget defined by the passed functions including sticky headers. The first returns the data size in rows and columns, second parameter is a function that returns cell template objects that can be cached and the third is used to apply data at specified data location to the passed template CanvasObject. The row and column headers will stick to the leading and top edges of the table and contain "1-10" and "A-Z" formatted labels.


<div class="since">Since: <code>
2.4</code></div>

#### func (*Table) CreateRenderer

```go
func (t *Table) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new renderer for the table.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Table) Cursor

```go
func (t *Table) Cursor() desktop.Cursor
```

#### func (*Table) DragEnd

```go
func (t *Table) DragEnd()
```

#### func (*Table) Dragged

```go
func (t *Table) Dragged(e *fyne.DragEvent)
```

#### func (*Table) FocusGained

```go
func (t *Table) FocusGained()
```
FocusGained is called after this table has gained focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Table) FocusLost

```go
func (t *Table) FocusLost()
```
FocusLost is called after this Table has lost focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Table) MouseDown

```go
func (t *Table) MouseDown(e *desktop.MouseEvent)
```
MouseDown response to desktop mouse event

#### func (*Table) MouseIn

```go
func (t *Table) MouseIn(ev *desktop.MouseEvent)
```

#### func (*Table) MouseMoved

```go
func (t *Table) MouseMoved(ev *desktop.MouseEvent)
```

#### func (*Table) MouseOut

```go
func (t *Table) MouseOut()
```

#### func (*Table) MouseUp

```go
func (t *Table) MouseUp(*desktop.MouseEvent)
```
MouseUp response to desktop mouse event

#### func (*Table) RefreshItem

```go
func (t *Table) RefreshItem(id TableCellID)
```
RefreshItem refreshes a single item, specified by the item ID passed in.


<div class="since">Since: <code>
2.4</code></div>

#### func (*Table) ScrollTo

```go
func (t *Table) ScrollTo(id TableCellID)
```
ScrollTo will scroll to the given cell without changing the selection. Attempting to scroll beyond the limits of the table will scroll to the edge of the table instead.


<div class="since">Since: <code>
2.1</code></div>

#### func (*Table) ScrollToBottom

```go
func (t *Table) ScrollToBottom()
```
ScrollToBottom scrolls to the last row in the table


<div class="since">Since: <code>
2.1</code></div>

#### func (*Table) ScrollToLeading

```go
func (t *Table) ScrollToLeading()
```
ScrollToLeading scrolls horizontally to the leading edge of the table


<div class="since">Since: <code>
2.1</code></div>

#### func (*Table) ScrollToTop

```go
func (t *Table) ScrollToTop()
```
ScrollToTop scrolls to the first row in the table


<div class="since">Since: <code>
2.1</code></div>

#### func (*Table) ScrollToTrailing

```go
func (t *Table) ScrollToTrailing()
```
ScrollToTrailing scrolls horizontally to the trailing edge of the table


<div class="since">Since: <code>
2.1</code></div>

#### func (*Table) Select

```go
func (t *Table) Select(id TableCellID)
```
Select will mark the specified cell as selected.

#### func (*Table) SetColumnWidth

```go
func (t *Table) SetColumnWidth(id int, width float32)
```
SetColumnWidth supports changing the width of the specified column. Columns normally take the width of the template cell returned from the CreateCell callback. The width parameter uses the same units as a fyne.Size type and refers to the internal content width not including the divider size.


<div class="since">Since: <code>
1.4.1</code></div>

#### func (*Table) SetRowHeight

```go
func (t *Table) SetRowHeight(id int, height float32)
```
SetRowHeight supports changing the height of the specified row. Rows normally take the height of the template cell returned from the CreateCell callback. The height parameter uses the same units as a fyne.Size type and refers to the internal content height not including the divider size.


<div class="since">Since: <code>
2.3</code></div>

#### func (*Table) Tapped

```go
func (t *Table) Tapped(e *fyne.PointEvent)
```

#### func (*Table) TouchCancel

```go
func (t *Table) TouchCancel(*mobile.TouchEvent)
```
TouchCancel response to mobile touch event

#### func (*Table) TouchDown

```go
func (t *Table) TouchDown(e *mobile.TouchEvent)
```
TouchDown response to mobile touch event

#### func (*Table) TouchUp

```go
func (t *Table) TouchUp(*mobile.TouchEvent)
```
TouchUp response to mobile touch event

#### func (*Table) TypedKey

```go
func (t *Table) TypedKey(event *fyne.KeyEvent)
```
TypedKey is called if a key event happens while this Table is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Table) TypedRune

```go
func (t *Table) TypedRune(_ rune)
```
TypedRune is called if a text event happens while this Table is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Table) Unselect

```go
func (t *Table) Unselect(id TableCellID)
```
Unselect will mark the cell provided by id as unselected.

#### func (*Table) UnselectAll

```go
func (t *Table) UnselectAll()
```
UnselectAll will mark all cells as unselected.


<div class="since">Since: <code>
2.1</code></div>
