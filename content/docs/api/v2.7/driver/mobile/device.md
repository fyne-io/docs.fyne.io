---
tags: [api]
title: mobile.Device
slug: device

aliases:
- /api/v2.7/driver/mobile/device

package: fyne.io/fyne/v2/driver/mobile
---


---
```go
import "fyne.io/fyne/v2/driver/mobile"
```

#

###

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
