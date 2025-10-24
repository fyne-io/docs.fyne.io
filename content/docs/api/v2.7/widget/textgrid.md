---
tags: [api]
title: widget.TextGrid
slug: textgrid

aliases:
- /api/v2.7/widget/textgrid

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

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

###

```go
func NewTextGrid() *TextGrid
```
NewTextGrid creates a new empty TextGrid widget.

###

```go
func NewTextGridFromString(content string) *TextGrid
```
NewTextGridFromString creates a new TextGrid widget with the specified string content.

###

```go
func (t *TextGrid) Append(text string)
```
Append will add new lines to the end of this TextGrid. The first character will be at the beginning of a new line and any newline characters will split the text further.


<div class="since">Since: <code>
2.6</code></div>

###

```go
func (t *TextGrid) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (t *TextGrid) CursorLocationForPosition(p fyne.Position) (row, col int)
```
CursorLocationForPosition returns the location where a cursor would be if it was located in the cell under the requested position. If the grid is scrolled the position will refer to the visible offset and not the distance from the top left of the overall document.


<div class="since">Since: <code>
2.6</code></div>

###

```go
func (t *TextGrid) MinSize() fyne.Size
```
MinSize returns the smallest size this widget can shrink to

###

```go
func (t *TextGrid) PositionForCursorLocation(row, col int) fyne.Position
```
PositionForCursorLocation returns the relative position in this TextGrid for the cell at position row, col. If the grid has been scrolled this will be taken into account so that the position compared to top left will refer to the requested location.


<div class="since">Since: <code>
2.6</code></div>

###

```go
func (t *TextGrid) Resize(size fyne.Size)
```
Resize is called when this widget changes size. We should make sure that we refresh cells.

###

```go
func (t *TextGrid) Row(row int) TextGridRow
```
Row returns a copy of the content in a specified row as a TextGridRow. If the index is out of bounds it returns an empty row object.

###

```go
func (t *TextGrid) RowText(row int) string
```
RowText returns a string representation of the content at the row specified. If the index is out of bounds it returns an empty string.

###

```go
func (t *TextGrid) ScrollToBottom()
```
ScrollToBottom will scroll content to container bottom - to show latest info which end user just added


<div class="since">Since: <code>
2.7</code></div>

###

```go
func (t *TextGrid) ScrollToTop()
```
ScrollToTop will scroll content to container top


<div class="since">Since: <code>
2.7</code></div>

###

```go
func (t *TextGrid) SetCell(row, col int, cell TextGridCell)
```
SetCell sets a grid data to the cell at named row and column.

###

```go
func (t *TextGrid) SetRow(row int, content TextGridRow)
```
SetRow updates the specified row of the grid's contents using the specified content and style and then refreshes. If the row is beyond the end of the current buffer it will be expanded. Tab characters are not padded with spaces.

###

```go
func (t *TextGrid) SetRowStyle(row int, style TextGridStyle)
```
SetRowStyle sets a grid style to all the cells cell at the specified row. Any cells in this row with their own style will override this value when displayed.

###

```go
func (t *TextGrid) SetRune(row, col int, r rune)
```
SetRune sets a character to the cell at named row and column.

###

```go
func (t *TextGrid) SetStyle(row, col int, style TextGridStyle)
```
SetStyle sets a grid style to the cell at named row and column.

###

```go
func (t *TextGrid) SetStyleRange(startRow, startCol, endRow, endCol int, style TextGridStyle)
```
SetStyleRange sets a grid style to all the cells between the start row and column through to the end row and column.

###

```go
func (t *TextGrid) SetText(text string)
```
SetText updates the buffer of this textgrid to contain the specified text. New lines and columns will be added as required. Lines are separated by '\n'. The grid will use default text style and any previous content and style will be removed. Tab characters are padded with spaces to the next tab stop.

###

```go
func (t *TextGrid) Text() string
```
Text returns the contents of the buffer as a single string (with no style information). It reconstructs the lines by joining with a `\n` character. Tab characters have padded spaces removed.
