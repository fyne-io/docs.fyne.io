---
layout: page
tags: [api]
title: Fyne API "binding.DataList"
package: fyne.io/fyne/v2/data/binding
---

# binding.DataList
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
