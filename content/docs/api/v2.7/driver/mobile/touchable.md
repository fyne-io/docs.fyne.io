---
tags: [api]
title: mobile.Touchable
slug: touchable

aliases:
- /api/v2.7/driver/mobile/touchable

package: fyne.io/fyne/v2/driver/mobile
---


---
```go
import "fyne.io/fyne/v2/driver/mobile"
```

#

###

```go
type Touchable interface {
	TouchDown(*TouchEvent)
	TouchUp(*TouchEvent)
	TouchCancel(*TouchEvent)
}
```

Touchable represents mobile touch events that can be sent to CanvasObjects
