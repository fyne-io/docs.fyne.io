---
tags: [api]
title: binding.ExternalString
slug: externalstring

aliases:
- /api/v2.7/data/binding/externalstring

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type ExternalString = ExternalItem[string]
```

ExternalString supports binding a string value to an external value.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindString(v *string) ExternalString
```
BindString returns a new bindable value that controls the contents of the provided string variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
