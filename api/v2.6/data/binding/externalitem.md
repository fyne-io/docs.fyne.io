---
layout: page
tags: [api]
title: Fyne API "binding.ExternalItem"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalItem
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
