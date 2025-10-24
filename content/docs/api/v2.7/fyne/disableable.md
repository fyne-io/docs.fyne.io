---
tags: [api]
title: fyne.Disableable
slug: disableable

aliases:
- /api/v2.7//disableable

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Disableable interface {
	Enable()
	Disable()
	Disabled() bool
}
```

Disableable describes any [CanvasObject] that can be disabled. This is primarily used with objects that also implement the Tappable interface.
