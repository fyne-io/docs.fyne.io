---
tags: [api]
title: fyne.PointEvent
slug: pointevent

aliases:
- /api/v2/fyne/pointevent.html
- /api/v2.0//pointevent
- /api/v2.0//pointevent.html
- /api/v2.1//pointevent
- /api/v2.1//pointevent.html
- /api/v2.2//pointevent
- /api/v2.2//pointevent.html
- /api/v2.3//pointevent
- /api/v2.3//pointevent.html
- /api/v2.4//pointevent
- /api/v2.4//pointevent.html
- /api/v2.5//pointevent
- /api/v2.5//pointevent.html
- /api/v2.6//pointevent
- /api/v2.6//pointevent.html
- /api/v2.7//pointevent
- /api/v2.7//pointevent.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type PointEvent

```go
type PointEvent struct {
	AbsolutePosition Position // The absolute position of the event
	Position         Position // The relative position of the event
}
```

PointEvent describes a pointer input event. The position is relative to the top-left of the [CanvasObject] this is triggered on.
