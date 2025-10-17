---
layout: page
tags: [api]
title: Fyne API "embedded.TouchMoveEvent"
package: fyne.io/fyne/v2/driver/embedded
---

# embedded.TouchMoveEvent
---
```go
import "fyne.io/fyne/v2/driver/embedded"
```

## Usage

#### type TouchMoveEvent

```go
type TouchMoveEvent struct {
	Position fyne.Position
	ID       int
}
```

TouchMoveEvent is for indicating that an embedded device touch screen or pointing device was moved whilst being pressed.


<div class="since">Since: <code>
2.7</code></div>
