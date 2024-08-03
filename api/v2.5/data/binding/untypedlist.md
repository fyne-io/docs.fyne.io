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

	Append(value any) error
	Get() ([]any, error)
	GetValue(index int) (any, error)
	Prepend(value any) error
	Remove(value any) error
	Set(list []any) error
	SetValue(index int, value any) error
}
```

UntypedList supports binding a list of any values.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewUntypedList

```go
func NewUntypedList() UntypedList
```
NewUntypedList returns a bindable list of any values.


<div class="since">Since: <code>
2.1</code></div>
