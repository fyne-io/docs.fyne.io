---
title: ProgressBar

weight: 5060

aliases:
- /widget/progressbar.html
- /tour/widget/progressbar
---

The progress bar widget has two forms, the standard progress bar
shows the user which `Value` has been reached, from `Min` to
`Max`. The default min is `0.0` and the max defaults to `1.0`.
To use the default values just call [widget.NewProgressBar()](/api/v2/widget/progressbar/#func--newprogressbar).
After creating you can set the `Value` field.

To set up a custom range you can set `Min` and `Max` fields
manually. The label will always show the percentage completion.

The other form of progress widget is the infinite progress bar.
This version simply shows that some activity is ongoing by
moving a segment of the bar from left to right and back again.
You create this using [widget.NewProgressBarInfinite()](/api/v2/widget/progressbarinfinite/#func--newprogressbarinfinite) and
it will start animating as soon as it is shown.

```go
package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ProgressBar Widget")

	progress := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()

	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)
		}
	}()

	myWindow.SetContent(container.NewVBox(progress, infinite))
	myWindow.ShowAndRun()
}
```
