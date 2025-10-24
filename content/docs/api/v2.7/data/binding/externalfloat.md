---
tags: [api]
title: binding.ExternalFloat
slug: externalfloat

aliases:
- /api/v2.7/data/binding/externalfloat

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type ExternalFloat = ExternalItem[float64]
```

ExternalFloat supports binding a float64 value to an external value.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindFloat(v *float64) ExternalFloat
```
BindFloat returns a new bindable value that controls the contents of the provided float64 variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
