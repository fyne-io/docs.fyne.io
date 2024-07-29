---
layout: page
tags: [api]
title: Fyne API "driver.WindowsWindowContext"
package: fyne.io/fyne/v2/driver
---

# driver.WindowsWindowContext
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
