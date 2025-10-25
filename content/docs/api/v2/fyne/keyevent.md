---
tags: [api]
title: fyne.KeyEvent
slug: keyevent

aliases:
- /api/v2/fyne/keyevent.html
- /api/v2.0//keyevent
- /api/v2.0//keyevent.html
- /api/v2.1//keyevent
- /api/v2.1//keyevent.html
- /api/v2.2//keyevent
- /api/v2.2//keyevent.html
- /api/v2.3//keyevent
- /api/v2.3//keyevent.html
- /api/v2.4//keyevent
- /api/v2.4//keyevent.html
- /api/v2.5//keyevent
- /api/v2.5//keyevent.html
- /api/v2.6//keyevent
- /api/v2.6//keyevent.html
- /api/v2.7//keyevent
- /api/v2.7//keyevent.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type KeyEvent

```go
type KeyEvent struct {
	// Name describes the keyboard event that is consistent across platforms.
	Name KeyName
	// Physical is a platform specific field that reports the hardware information of physical keyboard events.
	Physical HardwareKey
}
```

KeyEvent describes a keyboard input event.
