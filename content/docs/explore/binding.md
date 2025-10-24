---
layout: page
title: Data Binding
slug: binding

weight: 2110
since: 2.0
---

Data binding was introduced in Fyne v2.0.0 and makes it easier to connect
many widgets to a data source that will update over time.
the `data/binding` package has many helpful bindings that can manage most standard
types that will be used in an application.
A data binding can be managed using the binding API (for example `NewString`)
or it can be connected to an external item of data like (`BindInt(*int)`).

Widgets that support binding typically have a `...WithData` constructor to
set up the binding when creating the widget. You can also call `Bind()` and
`Unbind()` to manage the data of an existing widget.
The following example shows how you can manage a `String` data item that
is bound to a simple `Label` widget.

```go
package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	str := binding.NewString()
	go func() {
		dots := "....."
		for i := 5; i >= 0; i-- {
			str.Set("Count down" + dots[:i])
			time.Sleep(time.Second)
		}
		str.Set("Blast off!")
	}()

	w.SetContent(widget.NewLabelWithData(str))
	w.ShowAndRun()
}
```

You can find out more in the [data binding](/binding/) section of this site.
