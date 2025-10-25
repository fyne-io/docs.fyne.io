---
tags: [api]
title: mobile.Device
slug: device

aliases:
- /api/v2/driver/mobile/device.html
- /api/v2.0/driver/mobile/device
- /api/v2.0/driver/mobile/device.html
- /api/v2.1/driver/mobile/device
- /api/v2.1/driver/mobile/device.html
- /api/v2.2/driver/mobile/device
- /api/v2.2/driver/mobile/device.html
- /api/v2.3/driver/mobile/device
- /api/v2.3/driver/mobile/device.html
- /api/v2.4/driver/mobile/device
- /api/v2.4/driver/mobile/device.html
- /api/v2.5/driver/mobile/device
- /api/v2.5/driver/mobile/device.html
- /api/v2.6/driver/mobile/device
- /api/v2.6/driver/mobile/device.html
- /api/v2.7/driver/mobile/device
- /api/v2.7/driver/mobile/device.html

package: fyne.io/fyne/v2/driver/mobile
---


---
```go
import "fyne.io/fyne/v2/driver/mobile"
```

## Usage

#### type Device

```go
type Device interface {
	// Request that the mobile device show the touch screen keyboard (standard layout)
	ShowVirtualKeyboard()
	// Request that the mobile device show the touch screen keyboard (custom layout)
	ShowVirtualKeyboardType(KeyboardType)
	// Request that the mobile device dismiss the touch screen keyboard
	HideVirtualKeyboard()
}
```

Device describes functionality only available on mobile
