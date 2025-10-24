---
tags: [api]
title: driver (package)
slug: (package)

aliases:
- /api/v2.7/driver

package: fyne.io/fyne/v2/driver
---


---
```go
import "fyne.io/fyne/v2/driver"
```


#

###

```go
func RunNative(fn func(any) error) error
```
RunNative provides a way to execute code within the platform-specific runtime context for various runtimes. On Android this provides the JVM pointers required to execute various NDK calls or use JNI APIs.


<div class="since">Since: <code>
2.3</code></div>

###

 * [AndroidContext](androidcontext.html)
 * [AndroidWindowContext](androidwindowcontext.html)
 * [MacWindowContext](macwindowcontext.html)
 * [NativeWindow](nativewindow.html)
 * [UnknownContext](unknowncontext.html)
 * [WaylandWindowContext](waylandwindowcontext.html)
 * [WindowsWindowContext](windowswindowcontext.html)
 * [X11WindowContext](x11windowcontext.html)
