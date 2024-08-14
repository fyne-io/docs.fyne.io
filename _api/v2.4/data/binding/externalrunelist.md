---
layout: page
tags: [api]
title: Fyne API "binding.ExternalRuneList"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalRuneList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalRuneList

```go
type ExternalRuneList interface {
	RuneList

	Reload() error
}
```

ExternalRuneList supports binding a list of rune values from an external variable.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindRuneList

```go
func BindRuneList(v *[]rune) ExternalRuneList
```
BindRuneList returns a bound list of rune values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
