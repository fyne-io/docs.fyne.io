---
tags: [api]
title: desktop.Cursorable
slug: cursorable

aliases:
- /api/v2.7/driver/desktop/cursorable

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

#

###

```go
type Cursorable interface {
	Cursor() Cursor
}
```

Cursorable describes any CanvasObject that needs a cursor change
