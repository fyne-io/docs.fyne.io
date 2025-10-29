---
tags: [api]
title: binding.ExternalItem
slug: externalitem

aliases:
- /api/data/binding/externalitem
- /api/data/binding/externalitem.html
- /api/v2.0/data/binding/externalitem
- /api/v2.0/data/binding/externalitem.html
- /api/v2.1/data/binding/externalitem
- /api/v2.1/data/binding/externalitem.html
- /api/v2.2/data/binding/externalitem
- /api/v2.2/data/binding/externalitem.html
- /api/v2.3/data/binding/externalitem
- /api/v2.3/data/binding/externalitem.html
- /api/v2.4/data/binding/externalitem
- /api/v2.4/data/binding/externalitem.html
- /api/v2.5/data/binding/externalitem
- /api/v2.5/data/binding/externalitem.html
- /api/v2.6/data/binding/externalitem
- /api/v2.6/data/binding/externalitem.html
- /api/v2.7/data/binding/externalitem
- /api/v2.7/data/binding/externalitem.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalItem

```go
type ExternalItem[T any] interface {
	Item[T]
	Reload() error
}
```

ExternalItem supports binding any external value of type T.


<div class="since">Since: <code>
2.6</code></div>

#### func  BindItem

```go
func BindItem[T any](val *T, comparator func(T, T) bool) ExternalItem[T]
```
BindItem returns a new bindable value that controls the contents of the provided variable of type T. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.6</code></div>
