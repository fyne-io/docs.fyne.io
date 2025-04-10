---
layout: page
tags: [api]
title: Fyne API "driver.X11WindowContext"
package: fyne.io/fyne/v2/driver
---

# driver.X11WindowContext
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
