---
title: Data Conversion
slug: conversion

weight: 7040

aliases:
- /tour/binding/conversion
---

So far we have used data binding where the type of
data matches the output type (for example `String` and
`Label` or `Entry`). Often it will be desirable to
present data that is not already in the correct format.
To do this the `binding` package provides a number
of helpful conversion functions.

Most commonly this will be used to convert different
types of data into strings for displaying in `Label`
or `Entry` widgets. See in the code how we can convert
a `Float` to `String` using `binding.FloatToString`.
The original value can be edited by moving the slider.
Each time the data changes it will run the conversion
code and update any connected widgets.

It is also possible to use format strings to add more
natural output for the user interface.
You can see that our `short` binding is also converting
a `Float` to `String` but by using a `WithFormat` helper
we can pass a format string (similar to the `fmt` package)
to provide a custom output.

```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Conversion")

	f := binding.NewFloat()
	str := binding.FloatToString(f)
	short := binding.FloatToStringWithFormat(f, "%0.0f%%")
	f.Set(25.0)

	w.SetContent(container.NewVBox(
		widget.NewSliderWithData(0, 100.0, f),
		widget.NewLabelWithData(str),
		widget.NewLabelWithData(short),
	))

	w.ShowAndRun()
}
```

Lastly in this section we will look at [list](/binding/list) data.
