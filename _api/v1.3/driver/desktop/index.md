---
layout: page
tags: [api]
title: Fyne API desktop
---

# desktop
---
```go
import "fyne.io/fyne/driver/desktop"
```


## Usage

```go
const (
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

	// KeyCapsLock represents the caps lock key, tapping once is the down event then again is the up
	KeyCapsLock fyne.KeyName = "CapsLock"
)
```

#### types

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
