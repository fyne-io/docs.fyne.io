---
tags: [api]
title: binding.ExternalUntyped
slug: externaluntyped

aliases:
- /api/v2.7/data/binding/externaluntyped

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type ExternalUntyped = ExternalItem[any]
```

ExternalUntyped supports binding a any value to an external value.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func BindUntyped(v any) ExternalUntyped
```
BindUntyped returns a bindable any value that is bound to an external type. The parameter must be a pointer to the type you wish to bind.


<div class="since">Since: <code>
2.1</code></div>
