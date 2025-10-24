---
tags: [api]
title: mobile.Keyboardable
slug: keyboardable

aliases:
- /api/v2.0/driver/mobile/keyboardable
- /api/v2.1/driver/mobile/keyboardable
- /api/v2.2/driver/mobile/keyboardable
- /api/v2.3/driver/mobile/keyboardable
- /api/v2.4/driver/mobile/keyboardable
- /api/v2.5/driver/mobile/keyboardable
- /api/v2.6/driver/mobile/keyboardable
- /api/v2.7/driver/mobile/keyboardable

package: fyne.io/fyne/v2/driver/mobile
---


---
```go
import "fyne.io/fyne/v2/driver/mobile"
```

## Usage

#### type Keyboardable

```go
type Keyboardable interface {
	fyne.Focusable

	Keyboard() KeyboardType
}
```

Keyboardable describes any CanvasObject that needs a keyboard
