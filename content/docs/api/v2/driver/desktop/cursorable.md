---
tags: [api]
title: desktop.Cursorable
slug: cursorable

aliases:
- /api/v2.0/driver/desktop/cursorable
- /api/v2.1/driver/desktop/cursorable
- /api/v2.2/driver/desktop/cursorable
- /api/v2.3/driver/desktop/cursorable
- /api/v2.4/driver/desktop/cursorable
- /api/v2.5/driver/desktop/cursorable
- /api/v2.6/driver/desktop/cursorable
- /api/v2.7/driver/desktop/cursorable

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type Cursorable

```go
type Cursorable interface {
	Cursor() Cursor
}
```

Cursorable describes any CanvasObject that needs a cursor change
