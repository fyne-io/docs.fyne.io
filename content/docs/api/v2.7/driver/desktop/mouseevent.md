---
tags: [api]
title: desktop.MouseEvent
slug: mouseevent

aliases:
- /api/v2.7/driver/desktop/mouseevent

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

#

###

```go
type MouseEvent struct {
	fyne.PointEvent
	Button   MouseButton
	Modifier fyne.KeyModifier
}
```

MouseEvent contains data relating to desktop mouse events
