---
tags: [api]
title: binding.Struct
slug: struct

aliases:
- /api/data/binding/struct
- /api/data/binding/struct.html
- /api/v2.0/data/binding/struct
- /api/v2.0/data/binding/struct.html
- /api/v2.1/data/binding/struct
- /api/v2.1/data/binding/struct.html
- /api/v2.2/data/binding/struct
- /api/v2.2/data/binding/struct.html
- /api/v2.3/data/binding/struct
- /api/v2.3/data/binding/struct.html
- /api/v2.4/data/binding/struct
- /api/v2.4/data/binding/struct.html
- /api/v2.5/data/binding/struct
- /api/v2.5/data/binding/struct.html
- /api/v2.6/data/binding/struct
- /api/v2.6/data/binding/struct.html
- /api/v2.7/data/binding/struct
- /api/v2.7/data/binding/struct.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Struct

```go
type Struct interface {
	DataMap
	GetValue(string) (any, error)
	SetValue(string, any) error
	Reload() error
}
```

Struct is the base interface for a bound struct type.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindStruct

```go
func BindStruct(i any) Struct
```
BindStruct creates a new map binding of string to any using the struct passed as data. The key for each item is a string representation of each exported field with the value set as an any. Only exported fields are included.


<div class="since">Since: <code>
2.0</code></div>
