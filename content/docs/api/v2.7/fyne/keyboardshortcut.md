---
tags: [api]
title: fyne.KeyboardShortcut
slug: keyboardshortcut

aliases:
- /api/v2.7//keyboardshortcut

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type KeyboardShortcut interface {
	Shortcut
	Key() KeyName
	Mod() KeyModifier
}
```

KeyboardShortcut describes a shortcut meant to be triggered by a keyboard action.
