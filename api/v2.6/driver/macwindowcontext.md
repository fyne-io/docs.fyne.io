---
layout: page
tags: [api]
title: Fyne API "driver.MacWindowContext"
package: fyne.io/fyne/v2/driver
---

# driver.MacWindowContext
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
