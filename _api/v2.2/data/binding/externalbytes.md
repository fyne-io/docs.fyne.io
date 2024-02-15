---
layout: page
tags: [api]
title: Fyne API "binding.ExternalBytes"
---

# binding.ExternalBytes
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalBytes

```go
type ExternalBytes interface {
	Bytes
	Reload() error
}
```

ExternalBytes supports binding a []byte value to an external value.


<div class="since">Since: <code>
2.2</code></div>

#### func  BindBytes

```go
func BindBytes(v *[]byte) ExternalBytes
```
BindBytes returns a new bindable value that controls the contents of the provided []byte variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.2</code></div>
