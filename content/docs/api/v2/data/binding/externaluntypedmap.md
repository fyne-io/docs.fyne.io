---
tags: [api]
title: binding.ExternalUntypedMap
slug: externaluntypedmap

aliases:
- /api/data/binding/externaluntypedmap
- /api/data/binding/externaluntypedmap.html
- /api/v2.0/data/binding/externaluntypedmap
- /api/v2.0/data/binding/externaluntypedmap.html
- /api/v2.1/data/binding/externaluntypedmap
- /api/v2.1/data/binding/externaluntypedmap.html
- /api/v2.2/data/binding/externaluntypedmap
- /api/v2.2/data/binding/externaluntypedmap.html
- /api/v2.3/data/binding/externaluntypedmap
- /api/v2.3/data/binding/externaluntypedmap.html
- /api/v2.4/data/binding/externaluntypedmap
- /api/v2.4/data/binding/externaluntypedmap.html
- /api/v2.5/data/binding/externaluntypedmap
- /api/v2.5/data/binding/externaluntypedmap.html
- /api/v2.6/data/binding/externaluntypedmap
- /api/v2.6/data/binding/externaluntypedmap.html
- /api/v2.7/data/binding/externaluntypedmap
- /api/v2.7/data/binding/externaluntypedmap.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalUntypedMap

```go
type ExternalUntypedMap interface {
	UntypedMap
	Reload() error
}
```

ExternalUntypedMap is a map data binding with all values untyped (any), connected to an external data source.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindUntypedMap

```go
func BindUntypedMap(d *map[string]any) ExternalUntypedMap
```
BindUntypedMap creates a new map binding of string to any based on the data passed. If your code changes the content of the map this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
