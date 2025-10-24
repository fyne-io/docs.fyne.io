---
tags: [api]
title: binding.Untyped
slug: untyped

aliases:
- /api/v2.0/data/binding/untyped
- /api/v2.1/data/binding/untyped
- /api/v2.2/data/binding/untyped
- /api/v2.3/data/binding/untyped
- /api/v2.4/data/binding/untyped
- /api/v2.5/data/binding/untyped
- /api/v2.6/data/binding/untyped
- /api/v2.7/data/binding/untyped

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Untyped

```go
type Untyped = Item[any]
```

Untyped supports binding an any value.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewUntyped

```go
func NewUntyped() Untyped
```
NewUntyped returns a bindable any value that is managed internally.


<div class="since">Since: <code>
2.1</code></div>
