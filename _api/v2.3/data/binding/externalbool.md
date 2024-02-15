---
layout: page
tags: [api]
title: Fyne API "binding.ExternalBool"
---

# binding.ExternalBool
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalBool

```go
type ExternalBool interface {
	Bool
	Reload() error
}
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
