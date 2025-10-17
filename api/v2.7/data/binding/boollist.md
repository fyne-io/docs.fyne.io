---
layout: page
tags: [api]
title: Fyne API "binding.BoolList"
package: fyne.io/fyne/v2/data/binding
---

# binding.BoolList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type BoolList

```go
type BoolList = List[bool]
```

BoolList supports binding a list of bool values.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindPreferenceBoolList

```go
func BindPreferenceBoolList(key string, p fyne.Preferences) BoolList
```
BindPreferenceBoolList returns a bound list of bool values that is managed by the application preferences. Changes to this value will be saved to application storage and when the app starts the previous values will be read.


<div class="since">Since: <code>
2.6</code></div>
