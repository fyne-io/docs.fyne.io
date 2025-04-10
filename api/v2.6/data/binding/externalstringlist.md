---
layout: page
tags: [api]
title: Fyne API "binding.ExternalStringList"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalStringList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalStringList

```go
type ExternalStringList interface {
	StringList

	Reload() error
}
```

ExternalStringList supports binding a list of string values from an external variable.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindStringList

```go
func BindStringList(v *[]string) ExternalStringList
```
BindStringList returns a bound list of string values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
