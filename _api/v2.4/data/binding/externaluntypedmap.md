---
layout: page
tags: [api]
title: Fyne API "binding.ExternalUntypedMap"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalUntypedMap
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

ExternalUntypedMap is a map data binding with all values untyped (interface{}), connected to an external data source.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindUntypedMap

```go
func BindUntypedMap(d *map[string]interface{}) ExternalUntypedMap
```
BindUntypedMap creates a new map binding of string to interface{} based on the data passed. If your code changes the content of the map this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
