---
tags: [api]
title: fyne.KeyEvent
slug: keyevent

aliases:
- /api/v2.7//keyevent

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type KeyEvent struct {
	// Name describes the keyboard event that is consistent across platforms.
	Name KeyName
	// Physical is a platform specific field that reports the hardware information of physical keyboard events.
	Physical HardwareKey
}
```

KeyEvent describes a keyboard input event.
