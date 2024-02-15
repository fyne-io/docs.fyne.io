---
layout: page
tags: [api]
title: Fyne API "binding.FloatList"
---

# binding.FloatList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type FloatList

```go
type FloatList interface {
	DataList

	Append(float64) error
	Get() ([]float64, error)
	GetValue(int) (float64, error)
	Prepend(float64) error
	Set([]float64) error
	SetValue(int, float64) error
}
```

FloatList supports binding a list of float64 values.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewFloatList

```go
func NewFloatList() FloatList
```
NewFloatList returns a bindable list of float64 values.


<div class="since">Since: <code>
2.0</code></div>
