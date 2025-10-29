---
tags: [api]
title: desktop.Hoverable
slug: hoverable

aliases:
- /api/driver/desktop/hoverable
- /api/driver/desktop/hoverable.html
- /api/v2.0/driver/desktop/hoverable
- /api/v2.0/driver/desktop/hoverable.html
- /api/v2.1/driver/desktop/hoverable
- /api/v2.1/driver/desktop/hoverable.html
- /api/v2.2/driver/desktop/hoverable
- /api/v2.2/driver/desktop/hoverable.html
- /api/v2.3/driver/desktop/hoverable
- /api/v2.3/driver/desktop/hoverable.html
- /api/v2.4/driver/desktop/hoverable
- /api/v2.4/driver/desktop/hoverable.html
- /api/v2.5/driver/desktop/hoverable
- /api/v2.5/driver/desktop/hoverable.html
- /api/v2.6/driver/desktop/hoverable
- /api/v2.6/driver/desktop/hoverable.html
- /api/v2.7/driver/desktop/hoverable
- /api/v2.7/driver/desktop/hoverable.html

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type Hoverable

```go
type Hoverable interface {
	// MouseIn is a hook that is called if the mouse pointer enters the element.
	MouseIn(*MouseEvent)
	// MouseMoved is a hook that is called if the mouse pointer moved over the element.
	MouseMoved(*MouseEvent)
	// MouseOut is a hook that is called if the mouse pointer leaves the element.
	MouseOut()
}
```

Hoverable is used when a canvas object wishes to know if a pointer device moves over it.
