---
layout: page
tags: [api]
title: Fyne API "desktop.Driver"
package: fyne.io/fyne/v2/driver/desktop
---

# desktop.Driver
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type Driver

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
