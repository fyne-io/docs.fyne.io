---
layout: page
tags: [api]
title: Fyne API "desktop"
---

# desktop
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

Package desktop provides desktop specific driver functionality.

## Usage

```go
const (
	// KeyNone represents no key
	KeyNone fyne.KeyName = ""
	// KeyShiftLeft represents the left shift key
	KeyShiftLeft fyne.KeyName = "LeftShift"
	// KeyShiftRight represents the right shift key
	KeyShiftRight fyne.KeyName = "RightShift"
	// KeyControlLeft represents the left control key
	KeyControlLeft fyne.KeyName = "LeftControl"
	// KeyControlRight represents the right control key
	KeyControlRight fyne.KeyName = "RightControl"
	// KeyAltLeft represents the left alt key
	KeyAltLeft fyne.KeyName = "LeftAlt"
	// KeyAltRight represents the right alt key
	KeyAltRight fyne.KeyName = "RightAlt"
	// KeySuperLeft represents the left "Windows" key (or "Command" key on macOS)
	KeySuperLeft fyne.KeyName = "LeftSuper"
	// KeySuperRight represents the right "Windows" key (or "Command" key on macOS)
	KeySuperRight fyne.KeyName = "RightSuper"
	// KeyMenu represents the left or right menu / application key
	KeyMenu fyne.KeyName = "Menu"
	// KeyPrintScreen represents the key used to cause a screen capture
	KeyPrintScreen fyne.KeyName = "PrintScreen"

	// KeyCapsLock represents the caps lock key, tapping once is the down event then again is the up
	KeyCapsLock fyne.KeyName = "CapsLock"
)
```

```go
const (
	// ShiftModifier represents a shift key being held
	//
	// Deprecated: Use fyne.KeyModifierShift instead.
	ShiftModifier = fyne.KeyModifierShift
	// ControlModifier represents the ctrl key being held
	//
	// Deprecated: Use fyne.KeyModifierControl instead.
	ControlModifier = fyne.KeyModifierControl
	// AltModifier represents either alt keys being held
	//
	// Deprecated: Use fyne.KeyModifierAlt instead.
	AltModifier = fyne.KeyModifierAlt
	// SuperModifier represents either super keys being held
	//
	// Deprecated: Use fyne.KeyModifierSuper instead.
	SuperModifier = fyne.KeyModifierSuper
)
```

#### types

 * [App](app.html)
 * [Canvas](canvas.html)
 * [Cursor](cursor.html)
 * [Cursorable](cursorable.html)
 * [CustomShortcut](customshortcut.html)
 * [Driver](driver.html)
 * [Hoverable](hoverable.html)
 * [Keyable](keyable.html)
 * [Modifier](modifier.html)
 * [MouseButton](mousebutton.html)
 * [MouseEvent](mouseevent.html)
 * [Mouseable](mouseable.html)
 * [StandardCursor](standardcursor.html)
