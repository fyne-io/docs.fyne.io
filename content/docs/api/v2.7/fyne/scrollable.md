---
tags: [api]
title: fyne.Scrollable
slug: scrollable

aliases:
- /api/v2.7//scrollable

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Scrollable interface {
	Scrolled(*ScrollEvent)
}
```

Scrollable describes any [CanvasObject] that can also be scrolled. This is mostly used to implement the widget.ScrollContainer.
