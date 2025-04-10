---
layout: page
tags: [api]
title: Fyne API "binding.UntypedTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.UntypedTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type UntypedTree

```go
type UntypedTree interface {
	DataTree

	Append(parent, id string, value any) error
	Get() (map[string][]string, map[string]any, error)
	GetValue(id string) (any, error)
	Prepend(parent, id string, value any) error
	Remove(id string) error
	Set(ids map[string][]string, values map[string]any) error
	SetValue(id string, value any) error
}
```

UntypedTree supports binding a tree of any values.


<div class="since">Since: <code>
2.5</code></div>

#### func  NewUntypedTree

```go
func NewUntypedTree() UntypedTree
```
NewUntypedTree returns a bindable tree of any values.


<div class="since">Since: <code>
2.5</code></div>
