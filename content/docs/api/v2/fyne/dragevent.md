---
tags: [api]
title: fyne.DragEvent
slug: dragevent

aliases:
- /api/v2.0//dragevent
- /api/v2.1//dragevent
- /api/v2.2//dragevent
- /api/v2.3//dragevent
- /api/v2.4//dragevent
- /api/v2.5//dragevent
- /api/v2.6//dragevent
- /api/v2.7//dragevent

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type DragEvent

```go
type DragEvent struct {
	PointEvent
	Dragged Delta
}
```

DragEvent defines the parameters of a pointer or other drag event. The DraggedX and DraggedY fields show how far the item was dragged since the last event.
