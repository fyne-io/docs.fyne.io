---
weight: 10001
title: "GUI Overview"
slug: guiintro

aliases:
- /guiintro
---

Graphical applications are often more complicated to create than web
based or command line applications. Fyne changes this by utilizing the great design of Go
to make building beautiful graphical applications simple and fast.

For a graphical application to work we need to create a window and ask
the app to run. Doing so will start the event handling code that responds
to user input and updates the screen as our code runs.

This example creates a new application with a single window with the 
title "Hello". Inside this window we place a single label containing
the text "Hello Fyne!".

```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	myWindow.SetContent(widget.NewLabel("Hello Fyne!"))

	myWindow.ShowAndRun()
}
```

Once the window content is set we show it and run the application
(`Window.ShowAndRun()` is a shortcut for `Window.Show()` && `App.Run()`).
After calling Run() or ShowAndRun() our application will run and the
function will return after the window has been closed.
