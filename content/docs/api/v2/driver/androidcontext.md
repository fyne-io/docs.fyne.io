---
tags: [api]
title: driver.AndroidContext
slug: androidcontext

aliases:
- /api/driver/androidcontext
- /api/driver/androidcontext.html
- /api/v2.0/driver/androidcontext
- /api/v2.0/driver/androidcontext.html
- /api/v2.1/driver/androidcontext
- /api/v2.1/driver/androidcontext.html
- /api/v2.2/driver/androidcontext
- /api/v2.2/driver/androidcontext.html
- /api/v2.3/driver/androidcontext
- /api/v2.3/driver/androidcontext.html
- /api/v2.4/driver/androidcontext
- /api/v2.4/driver/androidcontext.html
- /api/v2.5/driver/androidcontext
- /api/v2.5/driver/androidcontext.html
- /api/v2.6/driver/androidcontext
- /api/v2.6/driver/androidcontext.html
- /api/v2.7/driver/androidcontext
- /api/v2.7/driver/androidcontext.html

package: fyne.io/fyne/v2/driver
---


---
```go
import "fyne.io/fyne/v2/driver"
```

## Usage

#### type AndroidContext

```go
type AndroidContext struct {
	VM, Env, Ctx uintptr
}
```

AndroidContext is passed to the RunNative callback when it is executed on an Android device. The VM, Env and Ctx pointers are required to make various calls into JVM methods.


<div class="since">Since: <code>
2.3</code></div>
