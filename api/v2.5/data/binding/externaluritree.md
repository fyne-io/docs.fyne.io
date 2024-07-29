---
layout: page
tags: [api]
title: Fyne API "binding.ExternalURITree"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalURITree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalURITree

```go
type ExternalURITree interface {
	URITree

	Reload() error
}
```

ExternalURITree supports binding a tree of fyne.URI values from an external variable.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindURITree

```go
func BindURITree(ids *map[string][]string, v *map[string]fyne.URI) ExternalURITree
```
BindURITree returns a bound tree of fyne.URI values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>
