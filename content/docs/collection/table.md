---
title: Table

weight: 6020

aliases:
- /collection/table.html
- /tour/widget/table
- /widget/table
--- 

The `Table` collection widget is like the [List](/collection/list) widget (another of the toolkit's collection widgets) with a two-dimensional index.
Like `List` this is designed to help build really performant
interfaces when lots of data is being presented.
Because of this the widget is not created with all the data embedded, but instead calls out to the data source when needed.

The `Table` uses callback functions to ask for data when it is required.
There are 3 main callbacks, `Length`, `CreateCell` and `UpdateCell`. The Length callback (passed first) is the simplest,
it returns how many items are in the data to be presented, the two ints it returns represent the row and column count.
The other two relate to the content templates.

The `CreateCell` callback returns a new template object, just like list.
The difference being that `MinSize` will define the standard size of each cell, and the minimum size of the table (it shows at least one cell).
As previously the `UpdateCell` is called to apply data to a cell template. The index passed in is the same `(row, col)` int pair.

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
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}
```

For more info, for example on how to add headers to the table, see the [widget.Table API documentation](/api/v2.5/widget/table.html).
