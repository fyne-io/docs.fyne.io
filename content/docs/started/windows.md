---
title: Window Handling
slug: windows

weight: 1060

aliases:
- /started/windows.html
- /tour/basics/windows
---

Windows are created using `App.NewWindow()` and need to be shown using
the `Show()` function. The helper method `ShowAndRun()` on [fyne.Window](/api/v2/fyne/window/#type--window)
allows you to show your window and run the application at the same time.

By default a window will be the right size to show its content
by checking the `MinSize()` function (more on that in later examples).
You can set a larger size by calling the `Window.Resize()` method.
Into this is passed a [fyne.Size](/api/v2/fyne/size/#type--size) which contains a width and height using
device independent pixels (meaning that it will be the same across different
devices), for example to make a window square by default we could:

```go
	w.Resize(fyne.NewSize(100, 100))
```

Be aware that the desktop environment may have constraints that cause
windows to be smaller than requested. Mobile devices will typically
ignore this as they are only displayed at full-screen.

If you wish to show a second window you must only call the `Show()`
function. It can also be helpful to split `Window.Show()` from `App.Run()`
if you want to open multiple windows when your application starts.
The example below shows how to load two windows when starting.

```go
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.Show()

	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewLabel("More content"))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	a.Run()
}
```

The above application will exit when both windows are closed. If your app
is arranged so one window is main and the others are accessory views you
can set one window to be "master" so that the app exits if that window
is closed. To do this use the `SetMaster()` function on `Window`.

Windows can be created at any time, we could change the code above so that
the content of the second window (`w2`) is a button that opens a new
window. You could also load windows from more complex workflows, but be
careful because new windows will normally appear above the current active
content.

```go
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))
```

{{% youtube "nM54LsMyBQY" %}}
