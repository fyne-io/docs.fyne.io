---
tags: [api]
title: binding.IntList
slug: intlist

aliases:
- /api/data/binding/intlist
- /api/data/binding/intlist.html
- /api/v2.0/data/binding/intlist
- /api/v2.0/data/binding/intlist.html
- /api/v2.1/data/binding/intlist
- /api/v2.1/data/binding/intlist.html
- /api/v2.2/data/binding/intlist
- /api/v2.2/data/binding/intlist.html
- /api/v2.3/data/binding/intlist
- /api/v2.3/data/binding/intlist.html
- /api/v2.4/data/binding/intlist
- /api/v2.4/data/binding/intlist.html
- /api/v2.5/data/binding/intlist
- /api/v2.5/data/binding/intlist.html
- /api/v2.6/data/binding/intlist
- /api/v2.6/data/binding/intlist.html
- /api/v2.7/data/binding/intlist
- /api/v2.7/data/binding/intlist.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type IntList

```go
type IntList = List[int]
```

IntList supports binding a list of int values.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindPreferenceIntList

```go
func BindPreferenceIntList(key string, p fyne.Preferences) IntList
```
BindPreferenceIntList returns a bound list of int values that is managed by the application preferences. Changes to this value will be saved to application storage and when the app starts the previous values will be read.


<div class="since">Since: <code>
2.6</code></div>
