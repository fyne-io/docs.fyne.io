---
tags: [api]
title: binding.DataTree
slug: datatree

aliases:
- /api/v2/data/binding/datatree.html
- /api/v2.0/data/binding/datatree
- /api/v2.0/data/binding/datatree.html
- /api/v2.1/data/binding/datatree
- /api/v2.1/data/binding/datatree.html
- /api/v2.2/data/binding/datatree
- /api/v2.2/data/binding/datatree.html
- /api/v2.3/data/binding/datatree
- /api/v2.3/data/binding/datatree.html
- /api/v2.4/data/binding/datatree
- /api/v2.4/data/binding/datatree.html
- /api/v2.5/data/binding/datatree
- /api/v2.5/data/binding/datatree.html
- /api/v2.6/data/binding/datatree
- /api/v2.6/data/binding/datatree.html
- /api/v2.7/data/binding/datatree
- /api/v2.7/data/binding/datatree.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type DataTree

```go
type DataTree interface {
	DataItem
	GetItem(id string) (DataItem, error)
	ChildIDs(string) []string
}
```

DataTree is the base interface for all bindable data trees.


<div class="since">Since: <code>
2.4</code></div>
