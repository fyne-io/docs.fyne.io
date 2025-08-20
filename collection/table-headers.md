---
title: Table Headers

--- 

The `Table` collection widget optionally supports 
"sticky" row and/or column headers, which 
remain at the top or left of the container while the data scrolls. 

This support uses additional callbacks, `CreateHeader` and `UpdateHeader`
to manage a separate template widget which can be 
styled separately from the data cell widget.

```go
package main

import (
    "fmt"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
)

var dataTemplateText = "0123456789012345"
var sampData = []string{
    "Lorem ipsum dolo",
    "consectetur adip",
    "sed do eiusmod t",
    "Ut enim ad minim",
    "quis nostrud exe",
}
var tableCols = 3

func makeTableComponents() *widget.Table {
    table := widget.NewTable(
        // length: return #rows, #cols
        func() (int, int) {
            return 15, tableCols
        },

        // create cell template
        func() fyne.CanvasObject {
            tpl := widget.NewLabel(dataTemplateText)
            return tpl
        },

        // update cell
        func(id widget.TableCellID, cellTpl fyne.CanvasObject) {
            cell := cellTpl.(*widget.Label)
            cell.SetText(sampData[(tableCols*id.Row+id.Col)%len(sampData)])
            // no explicit .Refresh() needed because .SetText() did it.
        },
    )

    // Enable sticky row and/or column headers
    table.ShowHeaderRow = true
    table.ShowHeaderColumn = true

    table.CreateHeader = func() fyne.CanvasObject {
        return widget.NewLabelWithStyle("row xx header",
            fyne.TextAlignCenter,
            fyne.TextStyle{Bold: true, Underline: true})
    }

    table.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
        cell := template.(*widget.Label)

        if id.Row < 0 && id.Col < 0 {
            // the top left header cell {-1, -1} is never populated
            panic(fmt.Sprintf("didn't expect update with id %v", id))
        } else if id.Row < 0 { // {row: -1, col: x} Set column header for col x
            cell.SetText(fmt.Sprintf("-- Col %d Header --", id.Col))
        } else if id.Col < 0 { // {row: x, col: -1} Set row header for row x
            cell.SetText(fmt.Sprintf("Row %d Header", id.Row))
        }
    }
    return table
}

func main() {
    a := app.NewWithID("com.example.sample.table_headers")
    w := a.NewWindow("table_headers")
    table := makeTableComponents()
    w.SetContent(table)
    w.Resize(fyne.NewSize(430, 300))
    w.ShowAndRun()
}
```

Which renders as:
![Sample program for sticky headers](fixed-headers.png)

Refer to [`widget.Table`](/api/widget/table) API for additional details on how to implement headers.

