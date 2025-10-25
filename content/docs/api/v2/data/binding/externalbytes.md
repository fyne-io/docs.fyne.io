---
tags: [api]
title: binding.ExternalBytes
slug: externalbytes

aliases:
- /api/v2/data/binding/externalbytes.html
- /api/v2.0/data/binding/externalbytes
- /api/v2.0/data/binding/externalbytes.html
- /api/v2.1/data/binding/externalbytes
- /api/v2.1/data/binding/externalbytes.html
- /api/v2.2/data/binding/externalbytes
- /api/v2.2/data/binding/externalbytes.html
- /api/v2.3/data/binding/externalbytes
- /api/v2.3/data/binding/externalbytes.html
- /api/v2.4/data/binding/externalbytes
- /api/v2.4/data/binding/externalbytes.html
- /api/v2.5/data/binding/externalbytes
- /api/v2.5/data/binding/externalbytes.html
- /api/v2.6/data/binding/externalbytes
- /api/v2.6/data/binding/externalbytes.html
- /api/v2.7/data/binding/externalbytes
- /api/v2.7/data/binding/externalbytes.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalBytes

```go
type ExternalBytes = ExternalItem[[]byte]
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
