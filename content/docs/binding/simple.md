---
title: Binding Simple Widgets
slug: simple

weight: 7020

aliases:
- /tour/binding/simple
---

The simplest form of binding a widget is to pass it
a bound item as a value instead of a static value.
Many widgets provide a `WithData` constructor that will
accept a typed data binding item. To set up the binding
all you need to do is pass the right type in.

Although this may not look like much of a benefit in the
initial code you can then see how it ensures that the
displayed content is always up to date with the source
of the data.
You will notice how we did not need to call `Refresh()`
on the `Label` widget or even keep a reference to it
and yet it updates appropriately.

```go
package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Simple")

	str := binding.NewString()
	str.Set("Initial value")

	text := widget.NewLabelWithData(str)
	w.SetContent(text)

	go func() {
		time.Sleep(time.Second * 2)
		str.Set("A new string")
	}()

	w.ShowAndRun()
}
```

In the next step we look at how to set up widgets 
that edit values through [two way](/binding/twoway) binding.
