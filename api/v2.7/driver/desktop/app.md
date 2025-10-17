---
layout: page
tags: [api]
title: Fyne API "desktop.App"
package: fyne.io/fyne/v2/driver/desktop
---

# desktop.App
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type App

```go
type App interface {
	SetSystemTrayMenu(menu *fyne.Menu)
	// SetSystemTrayIcon sets the icon to be used in system tray.
	// If you pass a `ThemedResource` then any OS that adjusts look to match theme will adapt the icon.
	SetSystemTrayIcon(icon fyne.Resource)

	// SetSystemTrayWindow optionally sets a window that this system tray will help to control.
	// On systems that support it (Windows, macOS and most Linux) the window will be shown on left-tap.
	// If the window is decorated (a regular window) tapping will show it only, however for a splash window
	// (without window decorations) tapping when the window is visible will hide it.
	// If you also have a menu set this will be triggered with right-mouse tap.
	// Note that your menu should probably include a "Show Window" menu item for less-compliant Linux systems.
	//
	// Since: 2.7
	SetSystemTrayWindow(fyne.Window)
}
```

App defines the desktop specific extensions to a fyne.App.


<div class="since">Since: <code>
2.2</code></div>
