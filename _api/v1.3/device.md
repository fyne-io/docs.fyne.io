---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Device

```go
type Device interface {
	Orientation() DeviceOrientation
	IsMobile() bool
	HasKeyboard() bool
	SystemScaleForWindow(Window) float32

	// Deprecated: Use SystemScaleForWindow instead - system scale can vary depending on window placement
	SystemScale() float32
}
```

Device provides information about the devices the code is running on

#### func  CurrentDevice

```go
func CurrentDevice() Device
```
CurrentDevice returns the device information for the current hardware (via the driver)
