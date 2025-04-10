---
layout: page
tags: [api]
title: Fyne API "mobile.Keyboardable"
package: fyne.io/fyne/v2/driver/mobile
---

# mobile.Keyboardable
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
