---
layout: page
tags: [api]
title: Fyne API "binding.ExternalInt"
---

# binding.ExternalInt
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalInt

```go
type ExternalInt interface {
	Int
	Reload() error
}
```

ExternalInt supports binding a int value to an external value.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindInt

```go
func BindInt(v *int) ExternalInt
```
BindInt returns a new bindable value that controls the contents of the provided int variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
