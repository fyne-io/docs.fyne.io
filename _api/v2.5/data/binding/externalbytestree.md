---
layout: page
tags: [api]
title: Fyne API "binding.ExternalBytesTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalBytesTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalBytesTree

```go
type ExternalBytesTree interface {
	BytesTree

	Reload() error
}
```

ExternalBytesTree supports binding a tree of []byte values from an external variable.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindBytesTree

```go
func BindBytesTree(ids *map[string][]string, v *map[string][]byte) ExternalBytesTree
```
BindBytesTree returns a bound tree of []byte values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>
