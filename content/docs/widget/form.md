---
title: Form

weight: 5050

aliases:
- /widget/form.html
- /tour/widget/form
---

The form widget is used to lay out many input fields with
labels and optional cancel and submit buttons. In its most
bare form it aligns labels to the left of each input widget.
By setting OnCancel or OnSubmit the form will add a button
bar with the specified handlers called when appropriate.

The widget can be created with `widget.NewForm(...)` passing
a list of `widget.FormItem`s, or by using the
`&widget.Form{}` syntax illustrated in the example.
There is also a helpful `Form.Append(label, widget)` that
can be used for an alternative syntax.

In this example we create two entries, one of which is a
"multiline" (like HTML TextArea) to hold values.
There is an OnSubmit handler which prints the information
before closing the window (and therefore the application).

```go
package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Widget")

	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			log.Println("multiline:", textArea.Text)
			myWindow.Close()
		},
	}

	// we can also append items
	form.Append("Text", textArea)

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
```
