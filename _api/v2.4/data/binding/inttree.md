---
layout: page
tags: [api]
title: Fyne API "binding.IntTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.IntTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type IntTree

```go
type IntTree interface {
	DataTree

	Append(parent, id string, value int) error
	Get() (map[string][]string, map[string]int, error)
	GetValue(id string) (int, error)
	Prepend(parent, id string, value int) error
	Set(ids map[string][]string, values map[string]int) error
	SetValue(id string, value int) error
}
```

IntTree supports binding a tree of int values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewIntTree

```go
func NewIntTree() IntTree
```
NewIntTree returns a bindable tree of int values.


<div class="since">Since: <code>
2.4</code></div>
