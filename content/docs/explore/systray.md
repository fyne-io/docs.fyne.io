---
layout: page
title: System Tray Menu
slug: systray

weight: 2100

since: 2.2
---

## Adding a System Tray menu

Since the v2.2.0 release Fyne has built in support for a system tray menu.
This feature displays an icon on macOS, Windows and Linux computers and when tapped
will pop out a menu as specified by the app.

As this is a desktop specific feature we must first do a runtime check that
the app is running in desktop mode. To do this, and get a reference to the
desktop features, we do a Go type assertion:

```go
	if desk, ok := a.(desktop.App); ok {
...
	}
```

If the `ok` variable is true then we can set up a menu using the standard
Fyne menu API that you might have used in `Window.SetMainMenu` before.

```go
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				log.Println("Tapped show")
			}))
		desk.SetSystemTrayMenu(m)
```

With this code added to the setup of your application you can run the app and
see that it shows a Fyne icon in the system tray. When you tap it a menu
will appear containing "Show" and "Quit".

The default icon is the Fyne logo, you can either fix this using [app metadata](/started/metadata)
or by setting the app icon in `App.SetIcon` or for system tray directly
using `desk.SetSystemTrayIcon`.

## Manage window lifecycle

By default a Fyne app will exit when you close all windows and this may not be
what you want with a system tray app. To override the behaviour you can use
the `Window.SetCloseIntercept` feature to override what happens when a window is
closed. In the example below we hide the window instead of closing it by calling
`Window.Hide()`. Add this before you show the window for the first time.

```go
	w.SetCloseIntercept(func() {
		w.Hide()
	})
```

The benefit of hiding a window is that you can simply show it again using
`Window.Show()` which is much more efficient than creating a new window if the
same content is needed a second time.
We update the menu created earlier to show the window that was hidden above.

```go
	fyne.NewMenuItem("Show", func() {
		w.Show()
	})
```

## Complete app

That's all there is to setting up a system tray menu with Fyne!
The complete code for this tutorial is as follows.

```go
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("SysTray")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	w.SetContent(widget.NewLabel("Fyne System Tray"))
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.ShowAndRun()
}
```

{{% youtube "t0vnGbzoB3I" %}}

