---
tags: [api]
title: fyne.Shortcut
slug: shortcut

aliases:
- /api/v2/fyne/shortcut.html
- /api/v2.0//shortcut
- /api/v2.0//shortcut.html
- /api/v2.1//shortcut
- /api/v2.1//shortcut.html
- /api/v2.2//shortcut
- /api/v2.2//shortcut.html
- /api/v2.3//shortcut
- /api/v2.3//shortcut.html
- /api/v2.4//shortcut
- /api/v2.4//shortcut.html
- /api/v2.5//shortcut
- /api/v2.5//shortcut.html
- /api/v2.6//shortcut
- /api/v2.6//shortcut.html
- /api/v2.7//shortcut
- /api/v2.7//shortcut.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Shortcut

```go
type Shortcut interface {
	ShortcutName() string
}
```

Shortcut is the interface used to describe a shortcut action
