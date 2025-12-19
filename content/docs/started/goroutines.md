---
title: Using Goroutines
slug: goroutines

weight: 1070
---

Since Fyne v2.6.0 all Fyne events and callbacks happen on the same goroutine - this allows
for higher performance and much better thread safety than any previous version.

With this change applications that wish to use goroutines (to do background updates for example)
have a little more work to do to get the same safety and performance as internal widgets.

## Updating from a goroutine

Any time your app invokes Fyne APIs from a goroutine your code created, you should use the
`fyne.Do` function. This tells Fyne that you want to queue changes to the application interface.
For example with a `canvas.Text` called `output` you may want to update the time each second.
The key code there would be:

```go
	fyne.Do(func() {
		output.Text = time.Now().Format(time.TimeOnly)
		output.Refresh()
	})
```

Such an update would occur in a goroutine that may be started from your main func.
You can see a complete clock example app below:

```go
package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	output := canvas.NewText(time.Now().Format(time.TimeOnly), color.NRGBA{G: 0xff, A: 0xff})
	output.TextStyle.Monospace = true
	output.TextSize = 32
	w.SetContent(output)

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			fyne.Do(func() {
				output.Text = time.Now().Format(time.TimeOnly)
				output.Refresh()
			})
		}
	}()
	w.ShowAndRun()
}

```

## Waiting for the code to run

When calling `fyne.Do` the code will execute between when you request it and the next frame to draw - which could be up to 15ms later. Some times it is necessary to know when that process has finished. For example if you are updating an image buffer and do not want to start again until after it is processed for the next visual update.

To do this another function, `fyne.DoAndWait` is useful. This takes the same parameter as the `Do` version above, but it will not return until the update is complete. This means the next line in your app will happen after the user interface is updated - very useful if you want to avoid potential issues with accessing shared resources concurrently.

{{% youtube "choFGqOPzDI" %}}

