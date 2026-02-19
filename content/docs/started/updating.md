---
layout: page
title: Updating Content in your GUI
slug: updating

weight: 1050
---

Having completed the [hello world](/started/hello) tutorial or other
examples you will have created a basic user interface. In this page
we see how the content of a GUI can be updated from your code.

The first step is to assign the widget you want to update to a
variable. In the hello world tutorial we passed [widget.NewLabel](/api/v2/widget/label/#func--newlabel)
directly into `SetContent()`, to update it we change that to two
different lines, such as:

```go
	message := widget.NewLabel("Welcome")
	w.SetContent(message)
```

Once the content has been assigned to a variable we can call functions
like `SetText("new text")`. For our example we will set the
content of our label to the current time, with the help of
`Time.Format`.

```go
	formatted := time.Now().Format("Time: 03:04:05")
	message.SetText(formatted)
```

That is all we need to do to change content of a visible item (see below for the full code).
In the following example we put that code into a button tap - now each time the
button is tapped the time will be displayed in the Label widget.

## Full example

The full code is as follows:

```go
package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Update Time")

	message := widget.NewLabel("Welcome")
	button := widget.NewButton("Update", func() {
		formatted := time.Now().Format("Time: 03:04:05")
		message.SetText(formatted)
	})

	w.SetContent(container.NewVBox(message, button))
	w.ShowAndRun()
}
```

{{% youtube "h2ZOdTA3ew4" %}}
