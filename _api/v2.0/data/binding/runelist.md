---
layout: page
tags: [api]
title: Fyne API "binding.RuneList"
---

# binding.RuneList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type RuneList

```go
type RuneList interface {
	DataList

	Append(rune) error
	Get() ([]rune, error)
	GetValue(int) (rune, error)
	Prepend(rune) error
	Set([]rune) error
	SetValue(int, rune) error
}
```

RuneList supports binding a list of rune values.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewRuneList

```go
func NewRuneList() RuneList
```
NewRuneList returns a bindable list of rune values.


<div class="since">Since: <code>
2.0</code></div>
