---
layout: page
tags: [api]
title: Fyne API "driver"
package: fyne.io/fyne/v2/driver
---

# driver
---
```go
import "fyne.io/fyne/v2/driver"
```


## Usage

#### func  RunNative

```go
func RunNative(fn func(any) error) error
```
RunNative provides a way to execute code within the platform-specific runtime context for various runtimes. This is mostly useful for Android where the JVM provides functionality that is not accessible directly in CGo. The call for most platforms will just execute passing an `UnknownContext` and returning any error reported.


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
