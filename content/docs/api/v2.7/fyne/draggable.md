---
tags: [api]
title: fyne.Draggable
slug: draggable

aliases:
- /api/v2.7//draggable

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Draggable interface {
	Dragged(*DragEvent)
	DragEnd()
}
```

Draggable indicates that a [CanvasObject] can be dragged. This is used for any item that the user has indicated should be moved across the screen.
