---
title: List

weight: 6010

aliases:
- /collection/list.html
  - /tour/widget/list
  - /widget/list
---

The `List` collection widget is one of the toolkit's collection widgets.
These widgets are designed to help build really performant
interfaces when lots of data is being presented.
You can also see the [Table](/collection/table) and [Tree](/collection/tree) widgets which have a
similar API. Because of this design they are a little more
complicated to use.

The `List` uses callback functions to ask for data when it is required.
There are 3 main callbacks, `Length`, `CreateItem` and `UpdateItem`. The Length callback (passed first) is the simplest,
it returns how many items are in the data to be presented. The
others relate to templates - how graphical elements are
created, cached and re-used.

The `CreateItem` callback returns a new template object. This
will be re-used with real data when the widget is presented.
The `MinSize` of this object will influence the `List` minimum size.
Lastly `UpdateItem` is called to apply an item of data to a
cached template. Use this to set the content ready for display.

```go
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"a", "string", "list"}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Widget")

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}
```
