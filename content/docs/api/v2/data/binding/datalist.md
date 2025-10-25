---
tags: [api]
title: binding.DataList
slug: datalist

aliases:
- /api/v2/data/binding/datalist.html
- /api/v2.0/data/binding/datalist
- /api/v2.0/data/binding/datalist.html
- /api/v2.1/data/binding/datalist
- /api/v2.1/data/binding/datalist.html
- /api/v2.2/data/binding/datalist
- /api/v2.2/data/binding/datalist.html
- /api/v2.3/data/binding/datalist
- /api/v2.3/data/binding/datalist.html
- /api/v2.4/data/binding/datalist
- /api/v2.4/data/binding/datalist.html
- /api/v2.5/data/binding/datalist
- /api/v2.5/data/binding/datalist.html
- /api/v2.6/data/binding/datalist
- /api/v2.6/data/binding/datalist.html
- /api/v2.7/data/binding/datalist
- /api/v2.7/data/binding/datalist.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type DataList

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
