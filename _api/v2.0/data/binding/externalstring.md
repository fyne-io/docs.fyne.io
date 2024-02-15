---
layout: page
tags: [api]
title: Fyne API "binding.ExternalString"
---

# binding.ExternalString
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalString

```go
type ExternalString interface {
	String
	Reload() error
}
```

ExternalString supports binding a string value to an external value.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindString

```go
func BindString(v *string) ExternalString
```
BindString returns a new bindable value that controls the contents of the provided string variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
