---
tags: [api]
title: binding.Untyped
slug: untyped

aliases:
- /api/v2.7/data/binding/untyped

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type Untyped = Item[any]
```

Untyped supports binding an any value.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func NewUntyped() Untyped
```
NewUntyped returns a bindable any value that is managed internally.


<div class="since">Since: <code>
2.1</code></div>
