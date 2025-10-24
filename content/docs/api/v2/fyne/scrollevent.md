---
tags: [api]
title: fyne.ScrollEvent
slug: scrollevent

aliases:
- /api/v2.0//scrollevent
- /api/v2.1//scrollevent
- /api/v2.2//scrollevent
- /api/v2.3//scrollevent
- /api/v2.4//scrollevent
- /api/v2.5//scrollevent
- /api/v2.6//scrollevent
- /api/v2.7//scrollevent

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ScrollEvent

```go
type ScrollEvent struct {
	PointEvent
	Scrolled Delta
}
```

ScrollEvent defines the parameters of a pointer or other scroll event. The DeltaX and DeltaY represent how large the scroll was in two dimensions.
