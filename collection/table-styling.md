---
title: Table Styling

redirect_from:
  - /tour/widget/table
  - /widget/table
--- 

The `Table` collection widget allows customization of the appearance 
of individual cells or whole rows or columns of your table.

Size, color, alignment or bolding can be changed to achieve such effects as:

* negative numbers in red
* numeric columns right adjusted
* different column widths for different columns, or different heights for different rows


## styling cells

The kinds of styling choices available depend on the `fyne.CanvasObject` returned from the `CreateCell` callback.  The most common template is a `widget.Label`, but a `widget.RichText` or `container.Stack` containing multiple widgets can be used.  

To achieve the desired effect, styling attributes are applied to each cell in
the `UpdateCell` callback based on the content or row and column position it has in the table. You might set an individual cell containing a negative number to display in red, or set all the cells in a given column to display centered, for example.  

Note that most widgets have transparent background.  To display a cell with a particular background color, the template can be a `container.Stack` containing a `canvas.Rectangle` for the background and a `widget.Label` for the foreground.

When changing an objects's styling attributes in the `UpdateCell` callback:

* call the object's `Refresh` method.  
However, an explicit `Refresh` is not necessary if the only attribute changes are made via setter methods, such as `label.SetText()`.
* set a given styling attribute on *all* code paths through the callback.  
The `fyne.CanvasObject` passed in to `UpdateCell` is not guaranteed to be the one initialized in `CreateCell` nor the one last displayed in that row and column;
its attribute settings on entry  are unpredictable.

## sizing rows and columns

The default column width and row height for the `Table` are determined by the `MinWidth` of the template object returned from `CreateCell` callback.  `MinWidth` cannot be set directly, but is influenced by content of the template object.  So, for example a template like  
`widget.NewLabel("123456789012345\nfoo")`  
would render a table with columns that are 15 characters wide by 2 lines high.  

The size of a particular column or row can be changed by:  
`func (t *Table) SetColumnWidth(id int, width float32)` and  
`func (t *Table) SetRowHeight(id int, height float32)`, respectively.

## sample table styling

```go
package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func makeTableComponents() *widget.Table {

	table := widget.NewTable(
		// length: return #rows, #cols
		func() (int, int) {
			return 15, 3
		},

		// create cell template
		func() fyne.CanvasObject {
			return widget.NewLabel("template text")
		},

		// update cell
		func(id widget.TableCellID, cellTpl fyne.CanvasObject) {
			cell := cellTpl.(*widget.Label)

			// style each column differently
			switch id.Col {
			case 0:
				cell.SetText(fmt.Sprintf(
					"row %d, wider, left aligned", id.Row))
				cell.Alignment = fyne.TextAlignLeading
			case 1:
				cell.SetText("centered")
				cell.Alignment = fyne.TextAlignCenter
			case 2:
				cell.SetText("right aligned")
				cell.Alignment = fyne.TextAlignTrailing
			}

			// style alternating rows differently
			if id.Row%2 == 1 {
				cell.TextStyle = fyne.TextStyle{Bold: true}
				cell.Importance = widget.DangerImportance
			} else {
				cell.TextStyle = fyne.TextStyle{}
				cell.Importance = widget.MediumImportance
			}

			cell.Refresh() // refresh needed after styling changes
		},
	)

	// set per-column widths as needed
	table.SetColumnWidth(0, 190)

	return table
}
func main() {
	a := app.NewWithID("com.example.sample.table_style")
	w := a.NewWindow("table_style")

	table := makeTableComponents()

	w.SetContent(table)
	w.Resize(fyne.NewSize(430, 300))
	w.ShowAndRun()
}

```


Next, we'll discuss setting [table headers](table-headers).
