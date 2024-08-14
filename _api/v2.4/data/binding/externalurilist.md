---
layout: page
tags: [api]
title: Fyne API "binding.ExternalURIList"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalURIList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalURIList

```go
type ExternalURIList interface {
	URIList

	Reload() error
}
```

ExternalURIList supports binding a list of fyne.URI values from an external variable.


<div class="since">Since: <code>
2.1</code></div>

#### func  BindURIList

```go
func BindURIList(v *[]fyne.URI) ExternalURIList
```
BindURIList returns a bound list of fyne.URI values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.1</code></div>
