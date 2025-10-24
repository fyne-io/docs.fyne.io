---
tags: [api]
title: binding.ExternalRune
slug: externalrune

aliases:
- /api/v2.7/data/binding/externalrune

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type ExternalRune = ExternalItem[rune]
```

ExternalRune supports binding a rune value to an external value.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindRune(v *rune) ExternalRune
```
BindRune returns a new bindable value that controls the contents of the provided rune variable. If your code changes the content of the variable this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
