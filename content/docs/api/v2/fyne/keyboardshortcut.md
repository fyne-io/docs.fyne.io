---
tags: [api]
title: fyne.KeyboardShortcut
slug: keyboardshortcut

aliases:
- /api//keyboardshortcut
- /api//keyboardshortcut.html
- /api/v2.0//keyboardshortcut
- /api/v2.0//keyboardshortcut.html
- /api/v2.1//keyboardshortcut
- /api/v2.1//keyboardshortcut.html
- /api/v2.2//keyboardshortcut
- /api/v2.2//keyboardshortcut.html
- /api/v2.3//keyboardshortcut
- /api/v2.3//keyboardshortcut.html
- /api/v2.4//keyboardshortcut
- /api/v2.4//keyboardshortcut.html
- /api/v2.5//keyboardshortcut
- /api/v2.5//keyboardshortcut.html
- /api/v2.6//keyboardshortcut
- /api/v2.6//keyboardshortcut.html
- /api/v2.7//keyboardshortcut
- /api/v2.7//keyboardshortcut.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type KeyboardShortcut

```go
type KeyboardShortcut interface {
	Shortcut
	Key() KeyName
	Mod() KeyModifier
}
```

KeyboardShortcut describes a shortcut meant to be triggered by a keyboard action.
