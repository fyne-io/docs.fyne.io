---
layout: page
tags: [api]
title: Fyne API "widget.Table"
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

	Length       func() (int, int)
	CreateCell   func() fyne.CanvasObject
	UpdateCell   func(id TableCellID, template fyne.CanvasObject)
	OnSelected   func(id TableCellID)
	OnUnselected func(id TableCellID)
}
```

Table widget is a grid of items that can be scrolled and a cell selected. It's performance is provided by caching cell templates created with CreateCell and re-using them with UpdateCell. The size of the content rows/columns is returned by the Length callback.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewTable

```go
func NewTable(length func() (int, int), create func() fyne.CanvasObject, update func(TableCellID, fyne.CanvasObject)) *Table
```
NewTable returns a new performant table widget defined by the passed functions. The first returns the data size in rows and columns, second parameter is a function that returns cell template objects that can be cached and the third is used to apply data at specified data location to the passed template CanvasObject.


<div class="since">Since: <code>
1.4</code></div>

#### func (*Table) CreateRenderer

```go
func (t *Table) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new renderer for the table.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Table) Select

```go
func (t *Table) Select(id TableCellID)
```
Select will mark the specified cell as selected.

#### func (*Table) SetColumnWidth

```go
func (t *Table) SetColumnWidth(id int, width float32)
```
SetColumnWidth supports changing the width of the specified column. Columns normally take the width of the template cell returned from the CreateCell callback. The width parameter uses the same units as a fyne.Size type and refers to the internal content width not including any standard padding or divider size.


<div class="since">Since: <code>
1.4.1</code></div>

#### func (*Table) Unselect

```go
func (t *Table) Unselect(id TableCellID)
```
Unselect will mark the cell provided by id as unselected.
