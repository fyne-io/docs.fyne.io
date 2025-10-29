---
tags: [api]
title: embedded.TouchUpEvent
slug: touchupevent

aliases:
- /api/driver/embedded/touchupevent
- /api/driver/embedded/touchupevent.html
- /api/v2.0/driver/embedded/touchupevent
- /api/v2.0/driver/embedded/touchupevent.html
- /api/v2.1/driver/embedded/touchupevent
- /api/v2.1/driver/embedded/touchupevent.html
- /api/v2.2/driver/embedded/touchupevent
- /api/v2.2/driver/embedded/touchupevent.html
- /api/v2.3/driver/embedded/touchupevent
- /api/v2.3/driver/embedded/touchupevent.html
- /api/v2.4/driver/embedded/touchupevent
- /api/v2.4/driver/embedded/touchupevent.html
- /api/v2.5/driver/embedded/touchupevent
- /api/v2.5/driver/embedded/touchupevent.html
- /api/v2.6/driver/embedded/touchupevent
- /api/v2.6/driver/embedded/touchupevent.html
- /api/v2.7/driver/embedded/touchupevent
- /api/v2.7/driver/embedded/touchupevent.html

package: fyne.io/fyne/v2/driver/embedded
---


---
```go
import "fyne.io/fyne/v2/driver/embedded"
```

## Usage

#### type TouchUpEvent

```go
type TouchUpEvent struct {
	Position fyne.Position
	ID       int
}
```

TouchUpEvent is for indicating that an embedded device touch screen or pointing device was released.


<div class="since">Since: <code>
2.7</code></div>
