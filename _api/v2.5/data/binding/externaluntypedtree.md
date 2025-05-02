---
layout: page
tags: [api]
title: Fyne API "binding.ExternalUntypedTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalUntypedTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalUntypedTree

```go
type ExternalUntypedTree interface {
	UntypedTree

	Reload() error
}
```

ExternalUntypedTree supports binding a tree of any values from an external variable.


<div class="since">Since: <code>
2.5</code></div>

#### func  BindUntypedTree

```go
func BindUntypedTree(ids *map[string][]string, v *map[string]any) ExternalUntypedTree
```
BindUntypedTree returns a bound tree of any values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>
