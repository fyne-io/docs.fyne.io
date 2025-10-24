---
tags: [api]
title: embedded.KeyEvent
slug: keyevent

aliases:
- /api/v2.0/driver/embedded/keyevent
- /api/v2.1/driver/embedded/keyevent
- /api/v2.2/driver/embedded/keyevent
- /api/v2.3/driver/embedded/keyevent
- /api/v2.4/driver/embedded/keyevent
- /api/v2.5/driver/embedded/keyevent
- /api/v2.6/driver/embedded/keyevent
- /api/v2.7/driver/embedded/keyevent

package: fyne.io/fyne/v2/driver/embedded
---


---
```go
import "fyne.io/fyne/v2/driver/embedded"
```

## Usage

#### type KeyEvent

```go
type KeyEvent struct {
	Name      fyne.KeyName
	Direction KeyDirection
}
```

KeyEvent is an event from keyboard actions occurring in an embedded device keyboard.


<div class="since">Since: <code>
2.7</code></div>
