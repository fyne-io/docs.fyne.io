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
}
```

App defines the desktop specific extensions to a fyne.App.


<div class="since">Since: <code>
2.2</code></div>
