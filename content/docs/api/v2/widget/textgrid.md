---
tags: [api]
title: widget.TextGrid
slug: textgrid

aliases:
- /api/widget/textgrid
- /api/widget/textgrid.html
- /api/v2.0/widget/textgrid
- /api/v2.0/widget/textgrid.html
- /api/v2.1/widget/textgrid
- /api/v2.1/widget/textgrid.html
- /api/v2.2/widget/textgrid
- /api/v2.2/widget/textgrid.html
- /api/v2.3/widget/textgrid
- /api/v2.3/widget/textgrid.html
- /api/v2.4/widget/textgrid
- /api/v2.4/widget/textgrid.html
- /api/v2.5/widget/textgrid
- /api/v2.5/widget/textgrid.html
- /api/v2.6/widget/textgrid
- /api/v2.6/widget/textgrid.html
- /api/v2.7/widget/textgrid
- /api/v2.7/widget/textgrid.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type TextGrid

```go
type TextGrid struct {
	BaseWidget
	Rows []TextGridRow

	ShowLineNumbers bool
	ShowWhitespace  bool
	TabWidth        int // If set to 0 the fyne.DefaultTabWidth is used

	// Scroll can be used to turn off the scrolling of our TextGrid.
	//
	// Since: 2.6
	Scroll fyne.ScrollDirection
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

#### func (*TextGrid) Append

```go
func (t *TextGrid) Append(text string)
```
Append will add new lines to the end of this TextGrid. The first character will be at the beginning of a new line and any newline characters will split the text further.


<div class="since">Since: <code>
2.6</code></div>

#### func (*TextGrid) CreateRenderer

```go
func (t *TextGrid) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*TextGrid) CursorLocationForPosition

```go
func (t *TextGrid) CursorLocationForPosition(p fyne.Position) (row, col int)
```
CursorLocationForPosition returns the location where a cursor would be if it was located in the cell under the requested position. If the grid is scrolled the position will refer to the visible offset and not the distance from the top left of the overall document.


<div class="since">Since: <code>
2.6</code></div>

#### func (*TextGrid) MinSize

```go
func (t *TextGrid) MinSize() fyne.Size
```
MinSize returns the smallest size this widget can shrink to

#### func (*TextGrid) PositionForCursorLocation

```go
func (t *TextGrid) PositionForCursorLocation(row, col int) fyne.Position
```
PositionForCursorLocation returns the relative position in this TextGrid for the cell at position row, col. If the grid has been scrolled this will be taken into account so that the position compared to top left will refer to the requested location.


<div class="since">Since: <code>
2.6</code></div>

#### func (*TextGrid) Resize

```go
func (t *TextGrid) Resize(size fyne.Size)
```
Resize is called when this widget changes size. We should make sure that we refresh cells.

#### func (*TextGrid) Row

```go
func (t *TextGrid) Row(row int) TextGridRow
```
Row returns a copy of the content in a specified row as a TextGridRow. If the index is out of bounds it returns an empty row object.

#### func (*TextGrid) RowText

```go
func (t *TextGrid) RowText(row int) string
```
RowText returns a string representation of the content at the row specified. If the index is out of bounds it returns an empty string.

#### func (*TextGrid) ScrollToBottom

```go
func (t *TextGrid) ScrollToBottom()
```
ScrollToBottom will scroll content to container bottom - to show latest info which end user just added


<div class="since">Since: <code>
2.7</code></div>

#### func (*TextGrid) ScrollToTop

```go
func (t *TextGrid) ScrollToTop()
```
ScrollToTop will scroll content to container top


<div class="since">Since: <code>
2.7</code></div>

#### func (*TextGrid) SetCell

```go
func (t *TextGrid) SetCell(row, col int, cell TextGridCell)
```
SetCell sets a grid data to the cell at named row and column.

#### func (*TextGrid) SetRow

```go
func (t *TextGrid) SetRow(row int, content TextGridRow)
```
SetRow updates the specified row of the grid's contents using the specified content and style and then refreshes. If the row is beyond the end of the current buffer it will be expanded. Tab characters are not padded with spaces.

#### func (*TextGrid) SetRowStyle

```go
func (t *TextGrid) SetRowStyle(row int, style TextGridStyle)
```
SetRowStyle sets a grid style to all the cells cell at the specified row. Any cells in this row with their own style will override this value when displayed.

#### func (*TextGrid) SetRune

```go
func (t *TextGrid) SetRune(row, col int, r rune)
```
SetRune sets a character to the cell at named row and column.

#### func (*TextGrid) SetStyle

```go
func (t *TextGrid) SetStyle(row, col int, style TextGridStyle)
```
SetStyle sets a grid style to the cell at named row and column.

#### func (*TextGrid) SetStyleRange

```go
func (t *TextGrid) SetStyleRange(startRow, startCol, endRow, endCol int, style TextGridStyle)
```
SetStyleRange sets a grid style to all the cells between the start row and column through to the end row and column.

#### func (*TextGrid) SetText

```go
func (t *TextGrid) SetText(text string)
```
SetText updates the buffer of this textgrid to contain the specified text. New lines and columns will be added as required. Lines are separated by '\n'. The grid will use default text style and any previous content and style will be removed. Tab characters are padded with spaces to the next tab stop.

#### func (*TextGrid) Text

```go
func (t *TextGrid) Text() string
```
Text returns the contents of the buffer as a single string (with no style information). It reconstructs the lines by joining with a `\n` character. Tab characters have padded spaces removed.
