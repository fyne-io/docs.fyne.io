---
layout: page
tags: [api]
title: Fyne API "binding.ExternalBoolTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalBoolTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalBoolTree

```go
type ExternalBoolTree interface {
	BoolTree

	Reload() error
}
```

ExternalBoolTree supports binding a tree of bool values from an external variable.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindBoolTree

```go
func BindBoolTree(ids *map[string][]string, v *map[string]bool) ExternalBoolTree
```
BindBoolTree returns a bound tree of bool values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>
