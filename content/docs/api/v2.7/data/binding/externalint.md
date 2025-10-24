---
tags: [api]
title: binding.ExternalInt
slug: externalint

aliases:
- /api/v2.7/data/binding/externalint

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type ExternalInt = ExternalItem[int]
```

ExternalInt supports binding a int value to an external value.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindInt(v *int) ExternalInt
```
BindInt returns a new bindable value that controls the contents of the provided int variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
