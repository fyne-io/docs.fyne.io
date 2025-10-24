---
tags: [api]
title: mobile.Keyboardable
slug: keyboardable

aliases:
- /api/v2.7/driver/mobile/keyboardable

package: fyne.io/fyne/v2/driver/mobile
---


---
```go
import "fyne.io/fyne/v2/driver/mobile"
```

#

###

```go
type Keyboardable interface {
	fyne.Focusable

	Keyboard() KeyboardType
}
```

Keyboardable describes any CanvasObject that needs a keyboard
