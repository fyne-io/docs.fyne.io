---
title: Two-Way Binding
slug: twoway

weight: 7030

aliases:
- /binding/twoway.html
- /tour/binding/twoway
---

So far we have looked at data binding as a way of keeping
our user interface elements up to date. Far more common,
however, is the need to update values from the UI
widgets and keep the data up to date everywhere.
Thankfully the bindings provided in Fyne are "two-way"
which means that values can be pushed into them as well
as read out. The change in data will be communicated to
all connected code without any additional code.

To see this in action we can update the last test app
to display a `Label` and an `Entry` that are bound to
the same value. By setting this up you can see that
editing the value through the entry will update the
text in the label as well. This is all possible without
calling refresh or referencing the widgets in our code.

By moving your app to use data binding you can stop
saving pointers to all your widgets. By instead 
capturing your data as a set of bound values your
user interface can be completely separate code.
Cleaner to read and easier to manage.

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
	w := myApp.NewWindow("Two Way")

	str := binding.NewString()
	str.Set("Hi!")

	w.SetContent(container.NewVBox(
		widget.NewLabelWithData(str),
		widget.NewEntryWithData(str),
	))

	w.ShowAndRun()
}
```

Next we will look at how to add [conversions](/binding/conversion) in our data.
