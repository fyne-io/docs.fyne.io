---
tags: [api]
title: embedded.TouchUpEvent
slug: touchupevent

aliases:
- /api/v2.7/driver/embedded/touchupevent

package: fyne.io/fyne/v2/driver/embedded
---


---
```go
import "fyne.io/fyne/v2/driver/embedded"
```

#

###

```go
type TouchUpEvent struct {
	Position fyne.Position
	ID       int
}
```

TouchUpEvent is for indicating that an embedded device touch screen or pointing device was released.


<div class="since">Since: <code>
2.7</code></div>
