---
tags: [api]
title: fyne.PointEvent
slug: pointevent

aliases:
- /api/v2.7//pointevent

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type PointEvent struct {
	AbsolutePosition Position // The absolute position of the event
	Position         Position // The relative position of the event
}
```

PointEvent describes a pointer input event. The position is relative to the top-left of the [CanvasObject] this is triggered on.
