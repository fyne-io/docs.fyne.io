---
layout: page
tags: [api]
title: Fyne API "binding.UntypedList"
package: fyne.io/fyne/v2/data/binding
---

# binding.UntypedList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type UntypedList

```go
type UntypedList interface {
	DataList

	Append(value interface{}) error
	Get() ([]interface{}, error)
	GetValue(index int) (interface{}, error)
	Prepend(value interface{}) error
	Set(list []interface{}) error
	SetValue(index int, value interface{}) error
}
```

UntypedList supports binding a list of interface{} values.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewUntypedList

```go
func NewUntypedList() UntypedList
```
NewUntypedList returns a bindable list of interface{} values.


<div class="since">Since: <code>
2.1</code></div>
