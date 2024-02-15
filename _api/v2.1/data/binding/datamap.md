---
layout: page
tags: [api]
title: Fyne API "binding.DataMap"
---

# binding.DataMap
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
