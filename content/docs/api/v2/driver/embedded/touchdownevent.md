---
tags: [api]
title: embedded.TouchDownEvent
slug: touchdownevent

aliases:
- /api/driver/embedded/touchdownevent
- /api/driver/embedded/touchdownevent.html
- /api/v2.0/driver/embedded/touchdownevent
- /api/v2.0/driver/embedded/touchdownevent.html
- /api/v2.1/driver/embedded/touchdownevent
- /api/v2.1/driver/embedded/touchdownevent.html
- /api/v2.2/driver/embedded/touchdownevent
- /api/v2.2/driver/embedded/touchdownevent.html
- /api/v2.3/driver/embedded/touchdownevent
- /api/v2.3/driver/embedded/touchdownevent.html
- /api/v2.4/driver/embedded/touchdownevent
- /api/v2.4/driver/embedded/touchdownevent.html
- /api/v2.5/driver/embedded/touchdownevent
- /api/v2.5/driver/embedded/touchdownevent.html
- /api/v2.6/driver/embedded/touchdownevent
- /api/v2.6/driver/embedded/touchdownevent.html
- /api/v2.7/driver/embedded/touchdownevent
- /api/v2.7/driver/embedded/touchdownevent.html

package: fyne.io/fyne/v2/driver/embedded
---


---
```go
import "fyne.io/fyne/v2/driver/embedded"
```

## Usage

#### type TouchDownEvent

```go
type TouchDownEvent struct {
	Position fyne.Position
	ID       int
}
```

TouchDownEvent is for indicating that an embedded device touch screen or pointing device was pressed.


<div class="since">Since: <code>
2.7</code></div>
