---
title: Text

weight: 3020

aliases:
- /canvas/text.html
- /tour/canvas/text
---

[canvas.Text](/api/v2/canvas/text/#type--text) is used for all text rendering within Fyne.
It is created by specifying the text and colour for the text.
Text is rendered using the default font, specified by the current theme.

The text object allows certain configuration like the `Alignment`
and `TextStyle` field. as illustrated in the example here.
To use a monospaced font instead you can specify
`fyne.TextStyle{Monospace: true}`.

```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Text")

	text := canvas.NewText("Text Object", color.White)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true}
	w.SetContent(text)

	w.ShowAndRun()
}
```

It is possible to use an alternative font by specifying a `FYNE_FONT`
environment variable. Use this to set a `.ttf` file to use instead of
the one provided in the Fyne toolkit or the current theme.
