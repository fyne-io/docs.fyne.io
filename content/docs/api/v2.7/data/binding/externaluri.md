---
tags: [api]
title: binding.ExternalURI
slug: externaluri

aliases:
- /api/v2.7/data/binding/externaluri

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type ExternalURI = ExternalItem[fyne.URI]
```

ExternalURI supports binding a fyne.URI value to an external value.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func BindURI(v *fyne.URI) ExternalURI
```
BindURI returns a new bindable value that controls the contents of the provided fyne.URI variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.1</code></div>
