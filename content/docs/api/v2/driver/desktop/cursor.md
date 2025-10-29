---
tags: [api]
title: desktop.Cursor
slug: cursor

aliases:
- /api/driver/desktop/cursor
- /api/driver/desktop/cursor.html
- /api/v2.0/driver/desktop/cursor
- /api/v2.0/driver/desktop/cursor.html
- /api/v2.1/driver/desktop/cursor
- /api/v2.1/driver/desktop/cursor.html
- /api/v2.2/driver/desktop/cursor
- /api/v2.2/driver/desktop/cursor.html
- /api/v2.3/driver/desktop/cursor
- /api/v2.3/driver/desktop/cursor.html
- /api/v2.4/driver/desktop/cursor
- /api/v2.4/driver/desktop/cursor.html
- /api/v2.5/driver/desktop/cursor
- /api/v2.5/driver/desktop/cursor.html
- /api/v2.6/driver/desktop/cursor
- /api/v2.6/driver/desktop/cursor.html
- /api/v2.7/driver/desktop/cursor
- /api/v2.7/driver/desktop/cursor.html

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type Cursor

```go
type Cursor interface {
	// Image returns the image for the given cursor, or nil if none should be shown.
	// It also returns the x and y pixels that should act as the hot-spot (measured from top left corner).
	Image() (image.Image, int, int)
}
```

Cursor interface is used for objects that desire a specific cursor.


<div class="since">Since: <code>
2.0</code></div>
