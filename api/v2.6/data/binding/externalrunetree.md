---
layout: page
tags: [api]
title: Fyne API "binding.ExternalRuneTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalRuneTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalRuneTree

```go
type ExternalRuneTree interface {
	RuneTree

	Reload() error
}
```

ExternalRuneTree supports binding a tree of rune values from an external variable.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindRuneTree

```go
func BindRuneTree(ids *map[string][]string, v *map[string]rune) ExternalRuneTree
```
BindRuneTree returns a bound tree of rune values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>
