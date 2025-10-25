---
tags: [api]
title: desktop.Mouseable
slug: mouseable

aliases:
- /api/v2/driver/desktop/mouseable.html
- /api/v2.0/driver/desktop/mouseable
- /api/v2.0/driver/desktop/mouseable.html
- /api/v2.1/driver/desktop/mouseable
- /api/v2.1/driver/desktop/mouseable.html
- /api/v2.2/driver/desktop/mouseable
- /api/v2.2/driver/desktop/mouseable.html
- /api/v2.3/driver/desktop/mouseable
- /api/v2.3/driver/desktop/mouseable.html
- /api/v2.4/driver/desktop/mouseable
- /api/v2.4/driver/desktop/mouseable.html
- /api/v2.5/driver/desktop/mouseable
- /api/v2.5/driver/desktop/mouseable.html
- /api/v2.6/driver/desktop/mouseable
- /api/v2.6/driver/desktop/mouseable.html
- /api/v2.7/driver/desktop/mouseable
- /api/v2.7/driver/desktop/mouseable.html

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type Mouseable

```go
type Mouseable interface {
	MouseDown(*MouseEvent)
	MouseUp(*MouseEvent)
}
```

Mouseable represents desktop mouse events that can be sent to CanvasObjects
