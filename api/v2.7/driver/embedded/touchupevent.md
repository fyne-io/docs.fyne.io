---
layout: page
tags: [api]
title: Fyne API "embedded.TouchUpEvent"
package: fyne.io/fyne/v2/driver/embedded
---

# embedded.TouchUpEvent
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
