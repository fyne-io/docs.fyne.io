---
layout: page
tags: [api]
title: Fyne API "binding.Bytes"
package: fyne.io/fyne/v2/data/binding
---

# binding.Bytes
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Bytes

```go
type Bytes interface {
	DataItem
	Get() ([]byte, error)
	Set([]byte) error
}
```

Bytes supports binding a []byte value.


<div class="since">Since: <code>
2.2</code></div>

#### func  NewBytes

```go
func NewBytes() Bytes
```
NewBytes returns a bindable []byte value that is managed internally.


<div class="since">Since: <code>
2.2</code></div>
