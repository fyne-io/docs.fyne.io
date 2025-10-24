---
tags: [api]
title: desktop.Driver
slug: driver

aliases:
- /api/v2.7/driver/desktop/driver

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

#

###

```go
type Driver interface {
	// Create a new borderless window that is centered on screen
	CreateSplashWindow() fyne.Window

	// Gets the set of key modifiers that are currently active
	//
	// Since: 2.4
	CurrentKeyModifiers() fyne.KeyModifier
}
```

Driver represents the extended capabilities of a desktop driver
