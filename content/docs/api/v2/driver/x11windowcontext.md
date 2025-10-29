---
tags: [api]
title: driver.X11WindowContext
slug: x11windowcontext

aliases:
- /api/driver/x11windowcontext
- /api/driver/x11windowcontext.html
- /api/v2.0/driver/x11windowcontext
- /api/v2.0/driver/x11windowcontext.html
- /api/v2.1/driver/x11windowcontext
- /api/v2.1/driver/x11windowcontext.html
- /api/v2.2/driver/x11windowcontext
- /api/v2.2/driver/x11windowcontext.html
- /api/v2.3/driver/x11windowcontext
- /api/v2.3/driver/x11windowcontext.html
- /api/v2.4/driver/x11windowcontext
- /api/v2.4/driver/x11windowcontext.html
- /api/v2.5/driver/x11windowcontext
- /api/v2.5/driver/x11windowcontext.html
- /api/v2.6/driver/x11windowcontext
- /api/v2.6/driver/x11windowcontext.html
- /api/v2.7/driver/x11windowcontext
- /api/v2.7/driver/x11windowcontext.html

package: fyne.io/fyne/v2/driver
---


---
```go
import "fyne.io/fyne/v2/driver"
```

## Usage

#### type X11WindowContext

```go
type X11WindowContext struct {
	// WindowHandle is the window handle for the native X11 window.
	WindowHandle uintptr
}
```

X11WindowContext is passed to the NativeWindow.RunNative callback when it is executed on a device with the X11 windowing system.


<div class="since">Since: <code>
2.5</code></div>
