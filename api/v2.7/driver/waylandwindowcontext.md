---
layout: page
tags: [api]
title: Fyne API "driver.WaylandWindowContext"
package: fyne.io/fyne/v2/driver
---

# driver.WaylandWindowContext
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
