---
tags: [api]
title: driver.WindowsWindowContext
slug: windowswindowcontext

aliases:
- /api/v2.7/driver/windowswindowcontext

package: fyne.io/fyne/v2/driver
---


---
```go
import "fyne.io/fyne/v2/driver"
```

#

###

```go
type WindowsWindowContext struct {
	// HWND is the window handle for the native window.
	HWND uintptr
}
```

WindowsWindowContext is passed to the NativeWindow.RunNative callback when it is executed on a Microsoft Windows device.


<div class="since">Since: <code>
2.5</code></div>
