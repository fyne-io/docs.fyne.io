---
layout: page
tags: [api]
title: Fyne API "driver.NativeWindow"
package: fyne.io/fyne/v2/driver
---

# driver.NativeWindow
---
```go
import "fyne.io/fyne/v2/driver"
```

## Usage

#### type NativeWindow

```go
type NativeWindow interface {
	// RunNative  provides a way to execute code within the platform-specific runtime context for a window.
	// The context types are defined in the `driver` package and the specific context passed will differ by platform.
	RunNative(func(context any))
}
```

NativeWindow is an extension interface for `fyne.Window` that gives access to platform-native features of application windows.


<div class="since">Since: <code>
2.5</code></div>
