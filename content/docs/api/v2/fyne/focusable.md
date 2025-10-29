---
tags: [api]
title: fyne.Focusable
slug: focusable

aliases:
- /api//focusable
- /api//focusable.html
- /api/v2.0//focusable
- /api/v2.0//focusable.html
- /api/v2.1//focusable
- /api/v2.1//focusable.html
- /api/v2.2//focusable
- /api/v2.2//focusable.html
- /api/v2.3//focusable
- /api/v2.3//focusable.html
- /api/v2.4//focusable
- /api/v2.4//focusable.html
- /api/v2.5//focusable
- /api/v2.5//focusable.html
- /api/v2.6//focusable
- /api/v2.6//focusable.html
- /api/v2.7//focusable
- /api/v2.7//focusable.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Focusable

```go
type Focusable interface {
	// FocusGained is a hook called by the focus handling logic after this object gained the focus.
	FocusGained()
	// FocusLost is a hook called by the focus handling logic after this object lost the focus.
	FocusLost()

	// TypedRune is a hook called by the input handling logic on text input events if this object is focused.
	TypedRune(rune)
	// TypedKey is a hook called by the input handling logic on key events if this object is focused.
	TypedKey(*KeyEvent)
}
```

Focusable describes any [CanvasObject] that can respond to being focused. It will receive the FocusGained and FocusLost events appropriately. When focused it will also have TypedRune called as text is input and TypedKey called when other keys are pressed.

Note: You must not change canvas state (including overlays or focus) in FocusGained or FocusLost or you would end up with a dead-lock.
