---
title: Entry

weight: 5030

aliases:
- /widget/entry.html
- /tour/widget/entry
---

The entry widget is used for user input of simple text content.
An entry can be created with a simple `widget.NewEntry()`
constructing function. When you create the widget keep a
reference so that you can access its `Text` field later.
It is also possible to use the `OnChanged` callback function
to be notified every time the content changes.

Entry widgets can also have validation for verifying the text
input typed into it. This can be done by setting the `Validator`
field to a `fyne.StringValidator`. You can also set a `PlaceHolder`
text and also set the entry to `MultiLine` to accept more than one
line of text.

```go
package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
```

You can also create a password entry (where the content is
obscured) using the `NewPasswordEntry()` function.
