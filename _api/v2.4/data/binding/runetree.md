---
layout: page
tags: [api]
title: Fyne API "binding.RuneTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.RuneTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type RuneTree

```go
type RuneTree interface {
	DataTree

	Append(parent, id string, value rune) error
	Get() (map[string][]string, map[string]rune, error)
	GetValue(id string) (rune, error)
	Prepend(parent, id string, value rune) error
	Set(ids map[string][]string, values map[string]rune) error
	SetValue(id string, value rune) error
}
```

RuneTree supports binding a tree of rune values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewRuneTree

```go
func NewRuneTree() RuneTree
```
NewRuneTree returns a bindable tree of rune values.


<div class="since">Since: <code>
2.4</code></div>
