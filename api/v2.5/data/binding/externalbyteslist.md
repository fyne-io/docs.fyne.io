---
layout: page
tags: [api]
title: Fyne API "binding.ExternalBytesList"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalBytesList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalBytesList

```go
type ExternalBytesList interface {
	BytesList

	Reload() error
}
```

ExternalBytesList supports binding a list of []byte values from an external variable.


<div class="since">Since: <code>
2.2</code></div>

#### func  BindBytesList

```go
func BindBytesList(v *[][]byte) ExternalBytesList
```
BindBytesList returns a bound list of []byte values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.2</code></div>
