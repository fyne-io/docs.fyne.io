---
tags: [api]
title: driver.AndroidWindowContext
slug: androidwindowcontext

aliases:
- /api/v2.7/driver/androidwindowcontext

package: fyne.io/fyne/v2/driver
---


---
```go
import "fyne.io/fyne/v2/driver"
```

#

###

```go
type AndroidWindowContext struct {
	AndroidContext
	NativeWindow uintptr
}
```

AndroidWindowContext is passed to the NativeWindow.RunNative callback when it is executed on an Android device. The NativeWindow field is of type `*C.ANativeWindow`. The VM, Env and Ctx pointers are required to make various calls into JVM methods.


<div class="since">Since: <code>
2.5</code></div>
