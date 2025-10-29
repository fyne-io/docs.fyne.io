---
tags: [api]
title: binding.DataItem
slug: dataitem

aliases:
- /api/data/binding/dataitem
- /api/data/binding/dataitem.html
- /api/v2.0/data/binding/dataitem
- /api/v2.0/data/binding/dataitem.html
- /api/v2.1/data/binding/dataitem
- /api/v2.1/data/binding/dataitem.html
- /api/v2.2/data/binding/dataitem
- /api/v2.2/data/binding/dataitem.html
- /api/v2.3/data/binding/dataitem
- /api/v2.3/data/binding/dataitem.html
- /api/v2.4/data/binding/dataitem
- /api/v2.4/data/binding/dataitem.html
- /api/v2.5/data/binding/dataitem
- /api/v2.5/data/binding/dataitem.html
- /api/v2.6/data/binding/dataitem
- /api/v2.6/data/binding/dataitem.html
- /api/v2.7/data/binding/dataitem
- /api/v2.7/data/binding/dataitem.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type DataItem

```go
type DataItem interface {
	// AddListener attaches a new change listener to this DataItem.
	// Listeners are called each time the data inside this DataItem changes.
	// Additionally, the listener will be triggered upon successful connection to get the current value.
	AddListener(DataListener)
	// RemoveListener will detach the specified change listener from the DataItem.
	// Disconnected listener will no longer be triggered when changes occur.
	RemoveListener(DataListener)
}
```

DataItem is the base interface for all bindable data items. All APIs on bindable data items are safe to invoke directly fron any goroutine.


<div class="since">Since: <code>
2.0</code></div>
