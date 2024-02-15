---
layout: page
tags: [api]
title: Fyne API "binding.StringList"
---

# binding.StringList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type StringList

```go
type StringList interface {
	DataList

	Append(value string) error
	Get() ([]string, error)
	GetValue(index int) (string, error)
	Prepend(value string) error
	Set(list []string) error
	SetValue(index int, value string) error
}
```

StringList supports binding a list of string values.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewStringList

```go
func NewStringList() StringList
```
NewStringList returns a bindable list of string values.


<div class="since">Since: <code>
2.0</code></div>
