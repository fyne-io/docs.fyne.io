---
tags: [api]
title: fyne.Shortcutable
slug: shortcutable

aliases:
- /api//shortcutable
- /api//shortcutable.html
- /api/v2.0//shortcutable
- /api/v2.0//shortcutable.html
- /api/v2.1//shortcutable
- /api/v2.1//shortcutable.html
- /api/v2.2//shortcutable
- /api/v2.2//shortcutable.html
- /api/v2.3//shortcutable
- /api/v2.3//shortcutable.html
- /api/v2.4//shortcutable
- /api/v2.4//shortcutable.html
- /api/v2.5//shortcutable
- /api/v2.5//shortcutable.html
- /api/v2.6//shortcutable
- /api/v2.6//shortcutable.html
- /api/v2.7//shortcutable
- /api/v2.7//shortcutable.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Shortcutable

```go
type Shortcutable interface {
	TypedShortcut(Shortcut)
}
```

Shortcutable describes any [CanvasObject] that can respond to shortcut commands (quit, cut, copy, and paste).
