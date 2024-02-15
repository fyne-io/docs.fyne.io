---
layout: page
tags: [api]
title: Fyne API "driver"
---

# driver
---
```go
import "fyne.io/fyne/v2/driver"
```


## Usage

#### func  RunNative

```go
func RunNative(fn func(interface{}) error) error
```
RunNative provides a way to execute code within the platform-specific runtime context for various runtimes. On Android this provides the JVM pointers required to execute various NDK calls or use JNI APIs.


<div class="since">Since: <code>
2.3</code></div>

#### types

 * [AndroidContext](androidcontext.html)
 * [UnknownContext](unknowncontext.html)
