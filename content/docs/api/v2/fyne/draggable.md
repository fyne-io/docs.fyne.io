---
tags: [api]
title: fyne.Draggable
slug: draggable

aliases:
- /api//draggable
- /api//draggable.html
- /api/v2.0//draggable
- /api/v2.0//draggable.html
- /api/v2.1//draggable
- /api/v2.1//draggable.html
- /api/v2.2//draggable
- /api/v2.2//draggable.html
- /api/v2.3//draggable
- /api/v2.3//draggable.html
- /api/v2.4//draggable
- /api/v2.4//draggable.html
- /api/v2.5//draggable
- /api/v2.5//draggable.html
- /api/v2.6//draggable
- /api/v2.6//draggable.html
- /api/v2.7//draggable
- /api/v2.7//draggable.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Draggable

```go
type Draggable interface {
	Dragged(*DragEvent)
	DragEnd()
}
```

Draggable indicates that a [CanvasObject] can be dragged. This is used for any item that the user has indicated should be moved across the screen.
