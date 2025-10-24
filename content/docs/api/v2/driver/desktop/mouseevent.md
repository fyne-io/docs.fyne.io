---
tags: [api]
title: desktop.MouseEvent
slug: mouseevent

aliases:
- /api/v2.0/driver/desktop/mouseevent
- /api/v2.1/driver/desktop/mouseevent
- /api/v2.2/driver/desktop/mouseevent
- /api/v2.3/driver/desktop/mouseevent
- /api/v2.4/driver/desktop/mouseevent
- /api/v2.5/driver/desktop/mouseevent
- /api/v2.6/driver/desktop/mouseevent
- /api/v2.7/driver/desktop/mouseevent

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type MouseEvent

```go
type MouseEvent struct {
	fyne.PointEvent
	Button   MouseButton
	Modifier fyne.KeyModifier
}
```

MouseEvent contains data relating to desktop mouse events
