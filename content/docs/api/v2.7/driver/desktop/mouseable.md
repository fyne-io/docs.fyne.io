---
tags: [api]
title: desktop.Mouseable
slug: mouseable

aliases:
- /api/v2.7/driver/desktop/mouseable

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

#

###

```go
type Mouseable interface {
	MouseDown(*MouseEvent)
	MouseUp(*MouseEvent)
}
```

Mouseable represents desktop mouse events that can be sent to CanvasObjects
