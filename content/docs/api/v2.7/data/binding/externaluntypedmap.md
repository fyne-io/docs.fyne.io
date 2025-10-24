---
tags: [api]
title: binding.ExternalUntypedMap
slug: externaluntypedmap

aliases:
- /api/v2.7/data/binding/externaluntypedmap

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type ExternalUntypedMap interface {
	UntypedMap
	Reload() error
}
```

ExternalUntypedMap is a map data binding with all values untyped (any), connected to an external data source.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindUntypedMap(d *map[string]any) ExternalUntypedMap
```
BindUntypedMap creates a new map binding of string to any based on the data passed. If your code changes the content of the map this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
