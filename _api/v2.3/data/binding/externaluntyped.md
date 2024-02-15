---
layout: page
tags: [api]
title: Fyne API "binding.ExternalUntyped"
---

# binding.ExternalUntyped
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalUntyped

```go
type ExternalUntyped interface {
	Untyped
	Reload() error
}
```

ExternalUntyped supports binding a interface{} value to an external value.


<div class="since">Since: <code>
2.1</code></div>

#### func  BindUntyped

```go
func BindUntyped(v interface{}) ExternalUntyped
```
BindUntyped returns a bindable interface{} value that is bound to an external type. The parameter must be a pointer to the type you wish to bind.


<div class="since">Since: <code>
2.1</code></div>
