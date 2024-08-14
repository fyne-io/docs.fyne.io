---
layout: page
tags: [api]
title: Fyne API "binding.BoolList"
package: fyne.io/fyne/v2/data/binding
---

# binding.BoolList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type BoolList

```go
type BoolList interface {
	DataList

	Append(value bool) error
	Get() ([]bool, error)
	GetValue(index int) (bool, error)
	Prepend(value bool) error
	Set(list []bool) error
	SetValue(index int, value bool) error
}
```

BoolList supports binding a list of bool values.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewBoolList

```go
func NewBoolList() BoolList
```
NewBoolList returns a bindable list of bool values.


<div class="since">Since: <code>
2.0</code></div>
