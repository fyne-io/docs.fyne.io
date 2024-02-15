---
layout: page
tags: [api]
title: Fyne API "binding.DataItem"
---

# binding.DataItem
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
	// Additionally the listener will be triggered upon successful connection to get the current value.
	AddListener(DataListener)
	// RemoveListener will detach the specified change listener from the DataItem.
	// Disconnected listener will no longer be triggered when changes occur.
	RemoveListener(DataListener)
}
```

DataItem is the base interface for all bindable data items.


<div class="since">Since: <code>
2.0</code></div>
