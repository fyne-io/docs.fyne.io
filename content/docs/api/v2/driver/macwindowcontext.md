---
tags: [api]
title: driver.MacWindowContext
slug: macwindowcontext

aliases:
- /api/v2/driver/macwindowcontext.html
- /api/v2.0/driver/macwindowcontext
- /api/v2.0/driver/macwindowcontext.html
- /api/v2.1/driver/macwindowcontext
- /api/v2.1/driver/macwindowcontext.html
- /api/v2.2/driver/macwindowcontext
- /api/v2.2/driver/macwindowcontext.html
- /api/v2.3/driver/macwindowcontext
- /api/v2.3/driver/macwindowcontext.html
- /api/v2.4/driver/macwindowcontext
- /api/v2.4/driver/macwindowcontext.html
- /api/v2.5/driver/macwindowcontext
- /api/v2.5/driver/macwindowcontext.html
- /api/v2.6/driver/macwindowcontext
- /api/v2.6/driver/macwindowcontext.html
- /api/v2.7/driver/macwindowcontext
- /api/v2.7/driver/macwindowcontext.html

package: fyne.io/fyne/v2/driver
---


---
```go
import "fyne.io/fyne/v2/driver"
```

## Usage

#### type MacWindowContext

```go
type MacWindowContext struct {
	// NSWindow is the window handle for the native window.
	NSWindow uintptr
}
```

MacWindowContext is passed to the NativeWindow.RunNative callback when it is executed on a macOS device.


<div class="since">Since: <code>
2.5</code></div>
