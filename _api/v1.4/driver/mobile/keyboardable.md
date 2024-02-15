---
layout: page
tags: [api]
title: Fyne API "mobile.Keyboardable"
---

# mobile.Keyboardable
---
```go
import "fyne.io/fyne/driver/mobile"
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
