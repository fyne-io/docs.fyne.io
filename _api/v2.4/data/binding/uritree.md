---
layout: page
tags: [api]
title: Fyne API "binding.URITree"
package: fyne.io/fyne/v2/data/binding
---

# binding.URITree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type URITree

```go
type URITree interface {
	DataTree

	Append(parent, id string, value fyne.URI) error
	Get() (map[string][]string, map[string]fyne.URI, error)
	GetValue(id string) (fyne.URI, error)
	Prepend(parent, id string, value fyne.URI) error
	Set(ids map[string][]string, values map[string]fyne.URI) error
	SetValue(id string, value fyne.URI) error
}
```

URITree supports binding a tree of fyne.URI values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewURITree

```go
func NewURITree() URITree
```
NewURITree returns a bindable tree of fyne.URI values.


<div class="since">Since: <code>
2.4</code></div>
