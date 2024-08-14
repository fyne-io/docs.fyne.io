---
layout: page
tags: [api]
title: Fyne API "binding.ExternalIntList"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalIntList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalIntList

```go
type ExternalIntList interface {
	IntList

	Reload() error
}
```

ExternalIntList supports binding a list of int values from an external variable.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindIntList

```go
func BindIntList(v *[]int) ExternalIntList
```
BindIntList returns a bound list of int values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
