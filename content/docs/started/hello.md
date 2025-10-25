---
layout: page
title: Creating your first Fyne app
slug: hello 

weight: 1020

aliases:
- /started/hello.html
- /started/firstapp
---

Having completed the steps in the [getting started](/started/) document you're ready to build your first app. To illustrate the process we will build a simple hello world application.

A simple app starts by creating an app instance with app.New() and then opening a window with app.NewWindow(). Then a widget tree is defined that is set as the main content with SetContent() on a window. The app UI is then shown by calling ShowAndRun() on the window.

<div id="hello__function">
<div style="text-align: left" markdown="1">

```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
```
</div>
</div>

The code above can be built using the command `go build .` and then executed either by running the `hello` command or by double clicking the icon. You could also bypass the compiling step and just run the code directly using `go run .`.

Either approach will show a window that looks just like this:

<img src="/images/started/hello-dark.png" style="width: 120px; border: 1px solid #999; margin: 10pt" alt="Hello Window" />

If you prefer a light theme then just set the environment variable `FYNE_THEME=light` and you'll get:

<img src="/images/started/hello-light.png" style="width: 120px; border: 1px solid #999; margin: 10pt" alt="Hello Light Theme" />

That's all there is to getting started. To learn more you can read the full
[API documentation](http://developer.fyne.io/api/).

{{% youtube "S3T9l9QUa9I" %}}
