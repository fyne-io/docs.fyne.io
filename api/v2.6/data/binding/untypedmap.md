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
	Get() (map[string]any, error)
	GetValue(string) (any, error)
	Set(map[string]any) error
	SetValue(string, any) error
}
```

UntypedMap is a map data binding with all values Untyped (any).


<div class="since">Since: <code>
2.0</code></div>

#### func  NewUntypedMap

```go
func NewUntypedMap() UntypedMap
```
NewUntypedMap creates a new, empty map binding of string to any.


<div class="since">Since: <code>
2.0</code></div>
