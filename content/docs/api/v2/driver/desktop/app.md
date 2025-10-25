---
tags: [api]
title: desktop.App
slug: app

aliases:
- /api/v2/driver/desktop/app.html
- /api/v2.0/driver/desktop/app
- /api/v2.0/driver/desktop/app.html
- /api/v2.1/driver/desktop/app
- /api/v2.1/driver/desktop/app.html
- /api/v2.2/driver/desktop/app
- /api/v2.2/driver/desktop/app.html
- /api/v2.3/driver/desktop/app
- /api/v2.3/driver/desktop/app.html
- /api/v2.4/driver/desktop/app
- /api/v2.4/driver/desktop/app.html
- /api/v2.5/driver/desktop/app
- /api/v2.5/driver/desktop/app.html
- /api/v2.6/driver/desktop/app
- /api/v2.6/driver/desktop/app.html
- /api/v2.7/driver/desktop/app
- /api/v2.7/driver/desktop/app.html

package: fyne.io/fyne/v2/driver/desktop
---


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
