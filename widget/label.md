---
title: Label

redirect_from:
  - /widget/
  - /tour/widget/label
  - /tour/widget/
---

Widgets are the main components of a Fyne application GUI, they can be
used in any place that a basic `fyne.CanvasObject` can. They manage user
interactions and will always match the current theme.

The Label widget is the simplest of them - it presents text to the user.
Unlike `canvas.Text` it can handle some simple formatting (such as `\n`)
and wrapping (by setting the `Wrapping` field).
You can create a label by calling `widget.NewLabel("some text")`, the
result can be assigned to a variable or passed directly into a container.

```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Label Widget")

	content := widget.NewLabel("text")

	content.Selectable = true
	content.SetText("New Label Widget Text")

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
```
