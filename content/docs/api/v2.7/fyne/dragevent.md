---
tags: [api]
title: fyne.DragEvent
slug: dragevent

aliases:
- /api/v2.7//dragevent

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type DragEvent struct {
	PointEvent
	Dragged Delta
}
```

DragEvent defines the parameters of a pointer or other drag event. The DraggedX and DraggedY fields show how far the item was dragged since the last event.
