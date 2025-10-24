---
tags: [api]
title: desktop.Keyable
slug: keyable

aliases:
- /api/v2.0/driver/desktop/keyable
- /api/v2.1/driver/desktop/keyable
- /api/v2.2/driver/desktop/keyable
- /api/v2.3/driver/desktop/keyable
- /api/v2.4/driver/desktop/keyable
- /api/v2.5/driver/desktop/keyable
- /api/v2.6/driver/desktop/keyable
- /api/v2.7/driver/desktop/keyable

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type Keyable

```go
type Keyable interface {
	fyne.Focusable

	KeyDown(*fyne.KeyEvent)
	KeyUp(*fyne.KeyEvent)
}
```

Keyable describes any focusable canvas object that can accept desktop key events. This is the traditional key down and up event that is not applicable to all devices.
