---
layout: page
tags: [api]
title: Fyne API "binding.ExternalFloatTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalFloatTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalFloatTree

```go
type ExternalFloatTree interface {
	FloatTree

	Reload() error
}
```

ExternalFloatTree supports binding a tree of float64 values from an external variable.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindFloatTree

```go
func BindFloatTree(ids *map[string][]string, v *map[string]float64) ExternalFloatTree
```
BindFloatTree returns a bound tree of float64 values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>
