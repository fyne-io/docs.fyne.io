---
layout: page
tags: [api]
title: Fyne API "binding.ExternalStringTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalStringTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalStringTree

```go
type ExternalStringTree interface {
	StringTree

	Reload() error
}
```

ExternalStringTree supports binding a tree of string values from an external variable.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindStringTree

```go
func BindStringTree(ids *map[string][]string, v *map[string]string) ExternalStringTree
```
BindStringTree returns a bound tree of string values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>
