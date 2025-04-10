---
layout: page
tags: [api]
title: Fyne API "binding.ExternalIntTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalIntTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalIntTree

```go
type ExternalIntTree interface {
	IntTree

	Reload() error
}
```

ExternalIntTree supports binding a tree of int values from an external variable.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindIntTree

```go
func BindIntTree(ids *map[string][]string, v *map[string]int) ExternalIntTree
```
BindIntTree returns a bound tree of int values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>
