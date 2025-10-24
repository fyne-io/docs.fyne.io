---
tags: [api]
title: fyne.Shortcutable
slug: shortcutable

aliases:
- /api/v2.7//shortcutable

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Shortcutable interface {
	TypedShortcut(Shortcut)
}
```

Shortcutable describes any [CanvasObject] that can respond to shortcut commands (quit, cut, copy, and paste).
