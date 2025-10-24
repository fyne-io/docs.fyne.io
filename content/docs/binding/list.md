---
title: List Data
slug: list

weight: 7050

aliases:
- /tour/binding/list
---

To demonstrate how more complex types can be connected
we will look at the `List` widget and how data binding
can make it easier to use. Firstly we create a `StringList`
data binding, which is a list of `String` data type.
Once we have a data of list type we can connect this to
the standard `List` widget. To do so we use the
`widget.NewListWithData` constructor, much like other
widgets.

Comparing this code to the [list tour](/widget/list)
You will see 2 main changes, the first is that we pass
the data type as the first parameter instead of a length
callback function. The second change is the last parameter,
our `UpdateItem` callback. The revised version takes
a `binding.DataItem` value instead of `widget.ListIndexID`.
When using this callback structure we should `Bind`
to the template label widget instead of calling `SetText`.
This means that if any of the strings change in the
data source each affected row of the table will refresh.

```go
package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Data")

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})
	myWindow.SetContent(container.NewBorder(nil, add, nil, nil, list))
	myWindow.ShowAndRun()
}
```

In our demo code there is an "Append" `Button`, when
tapped it will append a new value to the data source.
Doing so will automatically trigger the data change
handlers and expand the `List` widget to display the
new data.
