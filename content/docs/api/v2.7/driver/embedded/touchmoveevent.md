---
tags: [api]
title: embedded.TouchMoveEvent
slug: touchmoveevent

aliases:
- /api/v2.7/driver/embedded/touchmoveevent

package: fyne.io/fyne/v2/driver/embedded
---


---
```go
import "fyne.io/fyne/v2/driver/embedded"
```

#

###

```go
type TouchMoveEvent struct {
	Position fyne.Position
	ID       int
}
```

TouchMoveEvent is for indicating that an embedded device touch screen or pointing device was moved whilst being pressed.


<div class="since">Since: <code>
2.7</code></div>
