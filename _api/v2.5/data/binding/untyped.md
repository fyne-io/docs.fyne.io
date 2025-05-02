---
layout: page
tags: [api]
title: Fyne API "binding.Untyped"
package: fyne.io/fyne/v2/data/binding
---

# binding.Untyped
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Untyped

```go
type Untyped interface {
	DataItem
	Get() (any, error)
	Set(any) error
}
```

Untyped supports binding a any value.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewUntyped

```go
func NewUntyped() Untyped
```
NewUntyped returns a bindable any value that is managed internally.


<div class="since">Since: <code>
2.1</code></div>
