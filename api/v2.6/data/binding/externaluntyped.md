---
layout: page
tags: [api]
title: Fyne API "binding.ExternalUntyped"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalUntyped
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalUntyped

```go
type ExternalUntyped = ExternalItem[any]
```

ExternalUntyped supports binding a any value to an external value.


<div class="since">Since: <code>
2.1</code></div>

#### func  BindUntyped

```go
func BindUntyped(v any) ExternalUntyped
```
BindUntyped returns a bindable any value that is bound to an external type. The parameter must be a pointer to the type you wish to bind.


<div class="since">Since: <code>
2.1</code></div>
