---
tags: [api]
title: binding.ExternalBool
slug: externalbool

aliases:
- /api/v2.0/data/binding/externalbool
- /api/v2.1/data/binding/externalbool
- /api/v2.2/data/binding/externalbool
- /api/v2.3/data/binding/externalbool
- /api/v2.4/data/binding/externalbool
- /api/v2.5/data/binding/externalbool
- /api/v2.6/data/binding/externalbool
- /api/v2.7/data/binding/externalbool

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalBool

```go
type ExternalBool = ExternalItem[bool]
```

ExternalBool supports binding a bool value to an external value.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindBool

```go
func BindBool(v *bool) ExternalBool
```
BindBool returns a new bindable value that controls the contents of the provided bool variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
