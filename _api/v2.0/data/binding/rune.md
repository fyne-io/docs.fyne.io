---
layout: page
tags: [api]
title: Fyne API "binding.Rune"
---

# binding.Rune
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Rune

```go
type Rune interface {
	DataItem
	Get() (rune, error)
	Set(rune) error
}
```

Rune supports binding a rune value.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewRune

```go
func NewRune() Rune
```
NewRune returns a bindable rune value that is managed internally.


<div class="since">Since: <code>
2.0</code></div>
