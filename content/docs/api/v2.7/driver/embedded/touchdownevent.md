---
tags: [api]
title: embedded.TouchDownEvent
slug: touchdownevent

aliases:
- /api/v2.7/driver/embedded/touchdownevent

package: fyne.io/fyne/v2/driver/embedded
---


---
```go
import "fyne.io/fyne/v2/driver/embedded"
```

#

###

```go
type TouchDownEvent struct {
	Position fyne.Position
	ID       int
}
```

TouchDownEvent is for indicating that an embedded device touch screen or pointing device was pressed.


<div class="since">Since: <code>
2.7</code></div>
