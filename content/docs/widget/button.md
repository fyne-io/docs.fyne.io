---
title: Button

weight: 5020

aliases:
- /widget/button.html
  - /tour/widget/button
---

The button widget can contain text, an icon or both, the constructor
functions are `widget.NewButton()` and `widget.NewButtonWithIcon()`.
To create a text button there are just 2 parameters, the `string` content
and a 0 parameter `func()` that will be called when the button is tapped.
See the example for how that can be created.

The button constructor with an icon includes an additional parameter
which is the `fyne.Resource` which contains the icon data.
The builtin icons within the `theme` package all adapt appropriately
to a change in theme. You can pass in your own image if it is loaded
as a resource - helpers such as `fyne.LoadResourceFromPath()` may assist,
though bundling resources is recommended where possible.

To create a button with only an icon you should pass "" as the label
parameter to `widget.NewButtonWithIcon()`.

```go
package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")

	content := widget.NewButton("click me", func() {
		log.Println("tapped")
	})

	//content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	//})

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
````
