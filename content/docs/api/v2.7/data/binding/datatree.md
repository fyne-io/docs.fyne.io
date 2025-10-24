---
tags: [api]
title: binding.DataTree
slug: datatree

aliases:
- /api/v2.7/data/binding/datatree

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

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
