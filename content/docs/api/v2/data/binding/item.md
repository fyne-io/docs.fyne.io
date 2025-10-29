---
tags: [api]
title: binding.Item
slug: item

aliases:
- /api/data/binding/item
- /api/data/binding/item.html
- /api/v2.0/data/binding/item
- /api/v2.0/data/binding/item.html
- /api/v2.1/data/binding/item
- /api/v2.1/data/binding/item.html
- /api/v2.2/data/binding/item
- /api/v2.2/data/binding/item.html
- /api/v2.3/data/binding/item
- /api/v2.3/data/binding/item.html
- /api/v2.4/data/binding/item
- /api/v2.4/data/binding/item.html
- /api/v2.5/data/binding/item
- /api/v2.5/data/binding/item.html
- /api/v2.6/data/binding/item
- /api/v2.6/data/binding/item.html
- /api/v2.7/data/binding/item
- /api/v2.7/data/binding/item.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Item

```go
type Item[T any] interface {
	DataItem
	Get() (T, error)
	Set(T) error
}
```

Item supports binding any type T generically.


<div class="since">Since: <code>
2.6</code></div>

#### func  NewItem

```go
func NewItem[T any](comparator func(T, T) bool) Item[T]
```
NewItem returns a bindable value of type T that is managed internally.


<div class="since">Since: <code>
2.6</code></div>
