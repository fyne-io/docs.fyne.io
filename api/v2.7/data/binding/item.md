---
layout: page
tags: [api]
title: Fyne API "binding.Item"
package: fyne.io/fyne/v2/data/binding
---

# binding.Item
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
