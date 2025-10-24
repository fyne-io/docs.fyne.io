---
tags: [api]
title: fyne.Device
slug: device

aliases:
- /api/v2.7//device

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Device interface {
	Orientation() DeviceOrientation
	IsMobile() bool
	IsBrowser() bool
	HasKeyboard() bool
	SystemScaleForWindow(Window) float32

	// Locale returns the information about this device's language and region.
	//
	// Since: 2.5
	Locale() Locale
}
```

Device provides information about the devices the code is running on

###

```go
func CurrentDevice() Device
```
CurrentDevice returns the device information for the current hardware (via the driver)
