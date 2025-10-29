---
tags: [api]
title: fyne.DoubleTappable
slug: doubletappable

aliases:
- /api//doubletappable
- /api//doubletappable.html
- /api/v2.0//doubletappable
- /api/v2.0//doubletappable.html
- /api/v2.1//doubletappable
- /api/v2.1//doubletappable.html
- /api/v2.2//doubletappable
- /api/v2.2//doubletappable.html
- /api/v2.3//doubletappable
- /api/v2.3//doubletappable.html
- /api/v2.4//doubletappable
- /api/v2.4//doubletappable.html
- /api/v2.5//doubletappable
- /api/v2.5//doubletappable.html
- /api/v2.6//doubletappable
- /api/v2.6//doubletappable.html
- /api/v2.7//doubletappable
- /api/v2.7//doubletappable.html

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
