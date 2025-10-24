---
tags: [api]
title: binding.UntypedMap
slug: untypedmap

aliases:
- /api/v2.7/data/binding/untypedmap

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type UntypedMap interface {
	DataMap
	Delete(string)
	Get() (map[string]any, error)
	GetValue(string) (any, error)
	Set(map[string]any) error
	SetValue(string, any) error
}
```

UntypedMap is a map data binding with all values Untyped (any).


<div class="since">Since: <code>
2.0</code></div>

###

```go
func NewUntypedMap() UntypedMap
```
NewUntypedMap creates a new, empty map binding of string to any.


<div class="since">Since: <code>
2.0</code></div>
