---
tags: [api]
title: binding.BoolList
slug: boollist

aliases:
- /api/data/binding/boollist
- /api/data/binding/boollist.html
- /api/v2.0/data/binding/boollist
- /api/v2.0/data/binding/boollist.html
- /api/v2.1/data/binding/boollist
- /api/v2.1/data/binding/boollist.html
- /api/v2.2/data/binding/boollist
- /api/v2.2/data/binding/boollist.html
- /api/v2.3/data/binding/boollist
- /api/v2.3/data/binding/boollist.html
- /api/v2.4/data/binding/boollist
- /api/v2.4/data/binding/boollist.html
- /api/v2.5/data/binding/boollist
- /api/v2.5/data/binding/boollist.html
- /api/v2.6/data/binding/boollist
- /api/v2.6/data/binding/boollist.html
- /api/v2.7/data/binding/boollist
- /api/v2.7/data/binding/boollist.html

package: fyne.io/fyne/v2/data/binding
---


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
