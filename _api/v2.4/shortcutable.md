---
layout: page
tags: [api]
title: Fyne API "fyne.Shortcutable"
package: fyne.io/fyne/v2
---

# fyne.Shortcutable
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

Shortcutable describes any CanvasObject that can respond to shortcut commands (quit, cut, copy, and paste).
