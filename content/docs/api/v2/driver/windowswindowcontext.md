---
tags: [api]
title: driver.WindowsWindowContext
slug: windowswindowcontext

aliases:
- /api/v2/driver/windowswindowcontext.html
- /api/v2.0/driver/windowswindowcontext
- /api/v2.0/driver/windowswindowcontext.html
- /api/v2.1/driver/windowswindowcontext
- /api/v2.1/driver/windowswindowcontext.html
- /api/v2.2/driver/windowswindowcontext
- /api/v2.2/driver/windowswindowcontext.html
- /api/v2.3/driver/windowswindowcontext
- /api/v2.3/driver/windowswindowcontext.html
- /api/v2.4/driver/windowswindowcontext
- /api/v2.4/driver/windowswindowcontext.html
- /api/v2.5/driver/windowswindowcontext
- /api/v2.5/driver/windowswindowcontext.html
- /api/v2.6/driver/windowswindowcontext
- /api/v2.6/driver/windowswindowcontext.html
- /api/v2.7/driver/windowswindowcontext
- /api/v2.7/driver/windowswindowcontext.html

package: fyne.io/fyne/v2/driver
---


---
```go
import "fyne.io/fyne/v2/driver"
```

## Usage

#### type WindowsWindowContext

```go
type WindowsWindowContext struct {
	// HWND is the window handle for the native window.
	HWND uintptr
}
```

WindowsWindowContext is passed to the NativeWindow.RunNative callback when it is executed on a Microsoft Windows device.


<div class="since">Since: <code>
2.5</code></div>
