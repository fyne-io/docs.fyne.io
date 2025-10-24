---
tags: [api]
title: fyne.DoubleTappable
slug: doubletappable

aliases:
- /api/v2.0//doubletappable
- /api/v2.1//doubletappable
- /api/v2.2//doubletappable
- /api/v2.3//doubletappable
- /api/v2.4//doubletappable
- /api/v2.5//doubletappable
- /api/v2.6//doubletappable
- /api/v2.7//doubletappable

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type DoubleTappable

```go
type DoubleTappable interface {
	DoubleTapped(*PointEvent)
}
```

DoubleTappable describes any [CanvasObject] that can also be double tapped.
