---
title: Choices

weight: 5040

aliases:
- /widget/choices.html
- /tour/widget/choices
---

There are various widgets available to present the user with
a choice, these include a checkbox, radio group and select popup.

The [widget.Check](/api/v2/widget/check/#type--check) provides a simple yes/no choice and is created
using a string label. Each of these widgets also takes a
"changed" `func(...)` where the parameter is of the appropriate
type. `widget.NewCheck(..)` therefore takes a `string` parameter for
the label and a `func(bool)` param for the change handler.
You can also use the `Checked` field to get the boolean value.

The radio widget is similar, but the first parameter is a
slice of `string`s that represents each of the options.
The change function expects a `string` parameter this time
to return the currently selected value. Call `widget.NewRadioGroup(...)`
to construct the radio group widget, you can use this reference
later to read the `Selected` field instead of using the change
callback.

The select widget is identical in the constructor signature
as the radio widget. Calling `widget.NewSelect(...)` will
instead show a button that displays a popup when tapped from
which the user can make a selection. This is better for long
lists of options.

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
	myWindow := myApp.NewWindow("Choice Widgets")

	check := widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to", value)
	})
	radio := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to", value)
	})
	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Select set to", value)
	})

	myWindow.SetContent(container.NewVBox(check, radio, combo))
	myWindow.ShowAndRun()
}
```
