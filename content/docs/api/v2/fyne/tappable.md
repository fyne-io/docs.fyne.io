---
tags: [api]
title: fyne.Tappable
slug: tappable

aliases:
- /api/v2.0//tappable
- /api/v2.1//tappable
- /api/v2.2//tappable
- /api/v2.3//tappable
- /api/v2.4//tappable
- /api/v2.5//tappable
- /api/v2.6//tappable
- /api/v2.7//tappable

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Tappable

```go
type Tappable interface {
	Tapped(*PointEvent)
}
```

Tappable describes any [CanvasObject] that can also be tapped. This should be implemented by buttons etc that wish to handle pointer interactions.
