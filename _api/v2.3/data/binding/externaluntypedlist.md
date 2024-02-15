---
layout: page
tags: [api]
title: Fyne API "binding.ExternalUntypedList"
---

# binding.ExternalUntypedList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalUntypedList

```go
type ExternalUntypedList interface {
	UntypedList

	Reload() error
}
```

ExternalUntypedList supports binding a list of interface{} values from an external variable.


<div class="since">Since: <code>
2.1</code></div>

#### func  BindUntypedList

```go
func BindUntypedList(v *[]interface{}) ExternalUntypedList
```
BindUntypedList returns a bound list of interface{} values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.1</code></div>
