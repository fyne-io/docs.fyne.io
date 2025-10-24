---
tags: [api]
title: binding.IntList
slug: intlist

aliases:
- /api/v2.7/data/binding/intlist

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type IntList = List[int]
```

IntList supports binding a list of int values.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindPreferenceIntList(key string, p fyne.Preferences) IntList
```
BindPreferenceIntList returns a bound list of int values that is managed by the application preferences. Changes to this value will be saved to application storage and when the app starts the previous values will be read.


<div class="since">Since: <code>
2.6</code></div>
