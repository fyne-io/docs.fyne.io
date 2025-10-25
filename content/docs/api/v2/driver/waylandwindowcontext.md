---
tags: [api]
title: driver.WaylandWindowContext
slug: waylandwindowcontext

aliases:
- /api/v2/driver/waylandwindowcontext.html
- /api/v2.0/driver/waylandwindowcontext
- /api/v2.0/driver/waylandwindowcontext.html
- /api/v2.1/driver/waylandwindowcontext
- /api/v2.1/driver/waylandwindowcontext.html
- /api/v2.2/driver/waylandwindowcontext
- /api/v2.2/driver/waylandwindowcontext.html
- /api/v2.3/driver/waylandwindowcontext
- /api/v2.3/driver/waylandwindowcontext.html
- /api/v2.4/driver/waylandwindowcontext
- /api/v2.4/driver/waylandwindowcontext.html
- /api/v2.5/driver/waylandwindowcontext
- /api/v2.5/driver/waylandwindowcontext.html
- /api/v2.6/driver/waylandwindowcontext
- /api/v2.6/driver/waylandwindowcontext.html
- /api/v2.7/driver/waylandwindowcontext
- /api/v2.7/driver/waylandwindowcontext.html

package: fyne.io/fyne/v2/driver
---


---
```go
import "fyne.io/fyne/v2/driver"
```

## Usage

#### type WaylandWindowContext

```go
type WaylandWindowContext struct {
	// WaylandSurface is the handle to the native Wayland surface.
	WaylandSurface uintptr
}
```

WaylandWindowContext is passed to the NativeWindow.RunNative callback when it is executed on a device with the Wayland windowing system.


<div class="since">Since: <code>
2.5</code></div>
