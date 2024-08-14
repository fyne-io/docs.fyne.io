---
layout: page
tags: [api]
title: Fyne API "binding.StringTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.StringTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type StringTree

```go
type StringTree interface {
	DataTree

	Append(parent, id string, value string) error
	Get() (map[string][]string, map[string]string, error)
	GetValue(id string) (string, error)
	Prepend(parent, id string, value string) error
	Set(ids map[string][]string, values map[string]string) error
	SetValue(id string, value string) error
}
```

StringTree supports binding a tree of string values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewStringTree

```go
func NewStringTree() StringTree
```
NewStringTree returns a bindable tree of string values.


<div class="since">Since: <code>
2.4</code></div>
