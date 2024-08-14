---
layout: page
tags: [api]
title: Fyne API "binding.FloatTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.FloatTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type FloatTree

```go
type FloatTree interface {
	DataTree

	Append(parent, id string, value float64) error
	Get() (map[string][]string, map[string]float64, error)
	GetValue(id string) (float64, error)
	Prepend(parent, id string, value float64) error
	Set(ids map[string][]string, values map[string]float64) error
	SetValue(id string, value float64) error
}
```

FloatTree supports binding a tree of float64 values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewFloatTree

```go
func NewFloatTree() FloatTree
```
NewFloatTree returns a bindable tree of float64 values.


<div class="since">Since: <code>
2.4</code></div>
