---
tags: [api]
title: binding.ExternalUntyped
slug: externaluntyped

aliases:
- /api/v2/data/binding/externaluntyped.html
- /api/v2.0/data/binding/externaluntyped
- /api/v2.0/data/binding/externaluntyped.html
- /api/v2.1/data/binding/externaluntyped
- /api/v2.1/data/binding/externaluntyped.html
- /api/v2.2/data/binding/externaluntyped
- /api/v2.2/data/binding/externaluntyped.html
- /api/v2.3/data/binding/externaluntyped
- /api/v2.3/data/binding/externaluntyped.html
- /api/v2.4/data/binding/externaluntyped
- /api/v2.4/data/binding/externaluntyped.html
- /api/v2.5/data/binding/externaluntyped
- /api/v2.5/data/binding/externaluntyped.html
- /api/v2.6/data/binding/externaluntyped
- /api/v2.6/data/binding/externaluntyped.html
- /api/v2.7/data/binding/externaluntyped
- /api/v2.7/data/binding/externaluntyped.html

package: fyne.io/fyne/v2/data/binding
---


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
