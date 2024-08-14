---
layout: page
tags: [api]
title: Fyne API "binding.ExternalURI"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalURI
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalURI

```go
type ExternalURI interface {
	URI
	Reload() error
}
```

ExternalURI supports binding a fyne.URI value to an external value.


<div class="since">Since: <code>
2.1</code></div>

#### func  BindURI

```go
func BindURI(v *fyne.URI) ExternalURI
```
BindURI returns a new bindable value that controls the contents of the provided fyne.URI variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.1</code></div>
