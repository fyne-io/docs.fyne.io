---
layout: page
tags: [api]
title: Fyne API "binding.BoolTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.BoolTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type BoolTree

```go
type BoolTree interface {
	DataTree

	Append(parent, id string, value bool) error
	Get() (map[string][]string, map[string]bool, error)
	GetValue(id string) (bool, error)
	Prepend(parent, id string, value bool) error
	Set(ids map[string][]string, values map[string]bool) error
	SetValue(id string, value bool) error
}
```

BoolTree supports binding a tree of bool values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewBoolTree

```go
func NewBoolTree() BoolTree
```
NewBoolTree returns a bindable tree of bool values.


<div class="since">Since: <code>
2.4</code></div>
