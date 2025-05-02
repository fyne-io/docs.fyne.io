---
layout: page
tags: [api]
title: Fyne API "fyne.Device"
package: fyne.io/fyne/v2
---

# fyne.Device
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Device

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

#### func  CurrentDevice

```go
func CurrentDevice() Device
```
CurrentDevice returns the device information for the current hardware (via the driver)
