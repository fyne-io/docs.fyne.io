---
layout: page
tags: [api]
title: Fyne API "binding.BytesTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.BytesTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type BytesTree

```go
type BytesTree interface {
	DataTree

	Append(parent, id string, value []byte) error
	Get() (map[string][]string, map[string][]byte, error)
	GetValue(id string) ([]byte, error)
	Prepend(parent, id string, value []byte) error
	Set(ids map[string][]string, values map[string][]byte) error
	SetValue(id string, value []byte) error
}
```

BytesTree supports binding a tree of []byte values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewBytesTree

```go
func NewBytesTree() BytesTree
```
NewBytesTree returns a bindable tree of []byte values.


<div class="since">Since: <code>
2.4</code></div>
