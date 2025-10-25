---
tags: [api]
title: binding.Bytes
slug: bytes

aliases:
- /api/v2/data/binding/bytes.html
- /api/v2.0/data/binding/bytes
- /api/v2.0/data/binding/bytes.html
- /api/v2.1/data/binding/bytes
- /api/v2.1/data/binding/bytes.html
- /api/v2.2/data/binding/bytes
- /api/v2.2/data/binding/bytes.html
- /api/v2.3/data/binding/bytes
- /api/v2.3/data/binding/bytes.html
- /api/v2.4/data/binding/bytes
- /api/v2.4/data/binding/bytes.html
- /api/v2.5/data/binding/bytes
- /api/v2.5/data/binding/bytes.html
- /api/v2.6/data/binding/bytes
- /api/v2.6/data/binding/bytes.html
- /api/v2.7/data/binding/bytes
- /api/v2.7/data/binding/bytes.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Bytes

```go
type Bytes = Item[[]byte]
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
