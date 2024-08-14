---
layout: page
tags: [api]
title: Fyne API "binding.UntypedMap"
package: fyne.io/fyne/v2/data/binding
---

# binding.UntypedMap
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type UntypedMap

```go
type UntypedMap interface {
	DataMap
	Delete(string)
	Get() (map[string]interface{}, error)
	GetValue(string) (interface{}, error)
	Set(map[string]interface{}) error
	SetValue(string, interface{}) error
}
```

UntypedMap is a map data binding with all values Untyped (interface{}).


<div class="since">Since: <code>
2.0</code></div>

#### func  NewUntypedMap

```go
func NewUntypedMap() UntypedMap
```
NewUntypedMap creates a new, empty map binding of string to interface{}.


<div class="since">Since: <code>
2.0</code></div>
