---
tags: [api]
title: driver.MacWindowContext
slug: macwindowcontext

aliases:
- /api/v2.7/driver/macwindowcontext

package: fyne.io/fyne/v2/driver
---


---
```go
import "fyne.io/fyne/v2/driver"
```

#

###

```go
type MacWindowContext struct {
	// NSWindow is the window handle for the native window.
	NSWindow uintptr
}
```

MacWindowContext is passed to the NativeWindow.RunNative callback when it is executed on a macOS device.


<div class="since">Since: <code>
2.5</code></div>
