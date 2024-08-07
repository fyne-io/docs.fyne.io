---
layout: page
tags: [api]
title: Fyne API "binding.Struct"
package: fyne.io/fyne/v2/data/binding
---

# binding.Struct
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
