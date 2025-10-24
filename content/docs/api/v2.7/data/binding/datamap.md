---
tags: [api]
title: binding.DataMap
slug: datamap

aliases:
- /api/v2.7/data/binding/datamap

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type DataMap interface {
	DataItem
	GetItem(string) (DataItem, error)
	Keys() []string
}
```

DataMap is the base interface for all bindable data maps.


<div class="since">Since: <code>
2.0</code></div>
