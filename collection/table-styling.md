---
title: Table Styling

--- 

The default styling and size of each cell in the table is determined by 
the template object returned by the `CreateCell` callback of the [`Table`](/api/widget/table) collection widget.

These can be overridden for a given cell by styling the cell widget in the `UpdateCell` callback.

The size of a particular row or column can be overridden with
`Table` methods `.SetRowHeight()` or `.SetColumnWidth()`.

```go
package main

import (
    "fmt"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
)

func makeTableComponents() *widget.Table {
    wide_text := "row %2d, wider, left aligned"

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
                cell.SetText("right aligned")
                cell.Alignment = fyne.TextAlignTrailing
            case 1:
                cell.SetText("centered")
                cell.Alignment = fyne.TextAlignCenter
            case 2:
                cell.SetText(fmt.Sprintf(
                    wide_text, id.Row))
                cell.Alignment = fyne.TextAlignLeading
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

    // set an individual column width to fit a given string and styling
    stdTextWidth := fyne.CurrentApp().Settings().Theme().Size(theme.SizeNameText)
    strSize := fyne.MeasureText(wide_text, stdTextWidth, fyne.TextStyle{Bold: true})
    table.SetColumnWidth(2, strSize.Width)

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

Which renders as:
![Sample program showing table styling](./table-styling.png)

Several things to note about the `UpdateCell` callback:
1. Invoke the cell's `.Refresh()` method.  
Most widget "setter" methods do this automatically (e.g `cell.SetText()` 
in the example), but if the setter is not the last operation in the 
callback, the callback must do it explicitly.
2. Update all the styling attributes in each invocation.  
The `fyne.CanvasObject` provided on entry to the callback
may be a newly-instantiated template with initial styling
*or* a cell retrieved from cache which might have had other styling.  
The `UpdateCell` callback should update *all* the styling attributes
appropriate to the cell index and value.

Note that the sample overrides the width of one column
to fit the longest string in that column. 
It uses [`fyne.MeasureText()`](/api/size) 
to calculate the size of the string in screen units,
which is useful calculation in many contexts.

---

Next, we'll discuss setting [table headers](table-headers).
