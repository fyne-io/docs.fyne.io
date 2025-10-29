---
tags: [api]
title: driver (package)
slug: (package)

aliases:
- /api/driver
- /api/driver.html
- /api/v2.0/driver
- /api/v2.0/driver.html
- /api/v2.1/driver
- /api/v2.1/driver.html
- /api/v2.2/driver
- /api/v2.2/driver.html
- /api/v2.3/driver
- /api/v2.3/driver.html
- /api/v2.4/driver
- /api/v2.4/driver.html
- /api/v2.5/driver
- /api/v2.5/driver.html
- /api/v2.6/driver
- /api/v2.6/driver.html
- /api/v2.7/driver
- /api/v2.7/driver.html

package: fyne.io/fyne/v2/driver
---


---
```go
import "fyne.io/fyne/v2/driver"
```


## Usage

#### func  RunNative

```go
func RunNative(fn func(any) error) error
```
RunNative provides a way to execute code within the platform-specific runtime context for various runtimes. On Android this provides the JVM pointers required to execute various NDK calls or use JNI APIs.


<div class="since">Since: <code>
2.3</code></div>

#### types

 * [AndroidContext](androidcontext.html)
 * [AndroidWindowContext](androidwindowcontext.html)
 * [MacWindowContext](macwindowcontext.html)
 * [NativeWindow](nativewindow.html)
 * [UnknownContext](unknowncontext.html)
 * [WaylandWindowContext](waylandwindowcontext.html)
 * [WindowsWindowContext](windowswindowcontext.html)
 * [X11WindowContext](x11windowcontext.html)
