---
tags: [api]
title: desktop.Cursor
slug: cursor

aliases:
- /api/v2.7/driver/desktop/cursor

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

#

###

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
