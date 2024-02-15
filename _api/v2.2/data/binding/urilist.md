---
layout: page
tags: [api]
title: Fyne API "binding.URIList"
---

# binding.URIList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type URIList

```go
type URIList interface {
	DataList

	Append(value fyne.URI) error
	Get() ([]fyne.URI, error)
	GetValue(index int) (fyne.URI, error)
	Prepend(value fyne.URI) error
	Set(list []fyne.URI) error
	SetValue(index int, value fyne.URI) error
}
```

URIList supports binding a list of fyne.URI values.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewURIList

```go
func NewURIList() URIList
```
NewURIList returns a bindable list of fyne.URI values.


<div class="since">Since: <code>
2.1</code></div>
