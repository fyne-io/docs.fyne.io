---
tags: [api]
title: binding.DataList
slug: datalist

aliases:
- /api/v2.7/data/binding/datalist

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type DataList interface {
	DataItem
	GetItem(index int) (DataItem, error)
	Length() int
}
```

DataList is the base interface for all bindable data lists.


<div class="since">Since: <code>
2.0</code></div>
