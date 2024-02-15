---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type TextGrid

```go
type TextGrid struct {
	BaseWidget
	Rows []TextGridRow

	ShowLineNumbers bool
	ShowWhitespace  bool
}
```

TextGrid is a monospaced grid of characters. This is designed to be used by a text editor, code preview or terminal emulator.

#### func  NewTextGrid

```go
func NewTextGrid() *TextGrid
```
NewTextGrid creates a new empty TextGrid widget.

#### func  NewTextGridFromString

```go
func NewTextGridFromString(content string) *TextGrid
```
NewTextGridFromString creates a new TextGrid widget with the specified string content.

#### func (*TextGrid) CreateRenderer

```go
func (t *TextGrid) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to it's renderer

#### func (*TextGrid) MinSize

```go
func (t *TextGrid) MinSize() fyne.Size
```
MinSize returns the smallest size this widget can shrink to

#### func (*TextGrid) Resize

```go
func (t *TextGrid) Resize(size fyne.Size)
```
Resize is called when this widget changes size. We should make sure that we refresh cells.

#### func (*TextGrid) Row

```go
func (t *TextGrid) Row(row int) TextGridRow
```
Row returns the content of a specified row as a TextGridRow. If the index is out of bounds it returns an empty row object.

#### func (*TextGrid) SetRow

```go
func (t *TextGrid) SetRow(row int, content TextGridRow)
```
SetRow updates the specified row of the grid's contents using the specified content and style and then refreshes. If the row is beyond the end of the current buffer it will be expanded.

#### func (*TextGrid) SetStyle

```go
func (t *TextGrid) SetStyle(row, col int, style TextGridStyle)
```
SetStyle sets a grid style to the cell at named row and column

#### func (*TextGrid) SetStyleRange

```go
func (t *TextGrid) SetStyleRange(startRow, startCol, endRow, endCol int, style TextGridStyle)
```
SetStyleRange sets a grid style to all the cells between the start row and column through to the end row and column.

#### func (*TextGrid) SetText

```go
func (t *TextGrid) SetText(text string)
```
SetText updates the buffer of this textgrid to contain the specified text. New lines and columns will be added as required. Lines are separated by '\n'. The grid will use default text style and any previous content and style will be removed.

#### func (*TextGrid) Text

```go
func (t *TextGrid) Text() string
```
Text returns the contents of the buffer as a single string (with no style information). It reconstructs the lines by joining with a `\n` character.
