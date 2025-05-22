---
title: Table

redirect_from:
  - /tour/widget/table
  - /widget/table
--- 

The `Table` collection widget is like the [List](/collection/list) widget (another of the toolkit's collection widgets), but with a two-dimensional data structure.

It is designed to help build really performant
interfaces when lots of data is being presented.
The `Table` uses callback functions to ask for data when it is required.

* The `Length` callback returns the number of rows and columns of data from the data source.
The `Table` widget will display scrollbars if the layout doesn't have room to display all the data at once.

* The `CreateCell` callback returns an object which will be the template for all cells in the table.  
`MinSize` of the template object defines the standard size of each cell
and also the minimum size of the table (since it shows at least one cell).

* The `UpdateCell` callback applies data, content and styling for each cell of the table.  It is invoked for a particular cell only when that cell needs to be displayed, with a `TableCellID` specifying the row and column of the cell.

```go
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{[]string{"top left", "top right"},
	[]string{"bottom left", "bottom right"}}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")

	list := widget.NewTable(
		// Length callback
		func() (int, int) {	
			return len(data), len(data[0])
		},
		// CreateCell callback
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		// UpdateCell callback
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}
```

Next, we'll see how to [style](/collection/table-styling) the rows and columns of a table, 
and how to display [headers](/collection/table-headers).
