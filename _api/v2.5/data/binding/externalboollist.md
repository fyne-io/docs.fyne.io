---
layout: page
tags: [api]
title: Fyne API "binding.ExternalBoolList"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalBoolList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalBoolList

```go
type ExternalBoolList interface {
	BoolList

	Reload() error
}
```

ExternalBoolList supports binding a list of bool values from an external variable.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindBoolList

```go
func BindBoolList(v *[]bool) ExternalBoolList
```
BindBoolList returns a bound list of bool values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
