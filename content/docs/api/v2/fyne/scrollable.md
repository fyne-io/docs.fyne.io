---
tags: [api]
title: fyne.Scrollable
slug: scrollable

aliases:
- /api/v2.0//scrollable
- /api/v2.1//scrollable
- /api/v2.2//scrollable
- /api/v2.3//scrollable
- /api/v2.4//scrollable
- /api/v2.5//scrollable
- /api/v2.6//scrollable
- /api/v2.7//scrollable

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Scrollable

```go
type Scrollable interface {
	Scrolled(*ScrollEvent)
}
```

Scrollable describes any [CanvasObject] that can also be scrolled. This is mostly used to implement the widget.ScrollContainer.
