---
layout: page
tags: [api]
title: Fyne API "binding.ExternalRune"
---

# binding.ExternalRune
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalRune

```go
type ExternalRune interface {
	Rune
	Reload() error
}
```

ExternalRune supports binding a rune value to an external value.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindRune

```go
func BindRune(v *rune) ExternalRune
```
BindRune returns a new bindable value that controls the contents of the provided rune variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
