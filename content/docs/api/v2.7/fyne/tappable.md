---
tags: [api]
title: fyne.Tappable
slug: tappable

aliases:
- /api/v2.7//tappable

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Tappable interface {
	Tapped(*PointEvent)
}
```

Tappable describes any [CanvasObject] that can also be tapped. This should be implemented by buttons etc that wish to handle pointer interactions.
