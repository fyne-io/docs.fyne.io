---
tags: [api]
title: binding.DataMap
slug: datamap

aliases:
- /api/data/binding/datamap
- /api/data/binding/datamap.html
- /api/v2.0/data/binding/datamap
- /api/v2.0/data/binding/datamap.html
- /api/v2.1/data/binding/datamap
- /api/v2.1/data/binding/datamap.html
- /api/v2.2/data/binding/datamap
- /api/v2.2/data/binding/datamap.html
- /api/v2.3/data/binding/datamap
- /api/v2.3/data/binding/datamap.html
- /api/v2.4/data/binding/datamap
- /api/v2.4/data/binding/datamap.html
- /api/v2.5/data/binding/datamap
- /api/v2.5/data/binding/datamap.html
- /api/v2.6/data/binding/datamap
- /api/v2.6/data/binding/datamap.html
- /api/v2.7/data/binding/datamap
- /api/v2.7/data/binding/datamap.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type DataMap

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
