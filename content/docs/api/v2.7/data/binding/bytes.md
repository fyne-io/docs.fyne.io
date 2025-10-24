---
tags: [api]
title: binding.Bytes
slug: bytes

aliases:
- /api/v2.7/data/binding/bytes

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type Bytes = Item[[]byte]
```

Bytes supports binding a []byte value.


<div class="since">Since: <code>
2.2</code></div>

###

```go
func NewBytes() Bytes
```
NewBytes returns a bindable []byte value that is managed internally.


<div class="since">Since: <code>
2.2</code></div>
