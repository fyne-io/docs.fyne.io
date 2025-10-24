---
tags: [api]
title: fyne.DoubleTappable
slug: doubletappable

aliases:
- /api/v2.7//doubletappable

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type DoubleTappable interface {
	DoubleTapped(*PointEvent)
}
```

DoubleTappable describes any [CanvasObject] that can also be double tapped.
