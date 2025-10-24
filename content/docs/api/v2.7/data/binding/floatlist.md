---
tags: [api]
title: binding.FloatList
slug: floatlist

aliases:
- /api/v2.7/data/binding/floatlist

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type FloatList = List[float64]
```

FloatList supports binding a list of float64 values.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindPreferenceFloatList(key string, p fyne.Preferences) FloatList
```
BindPreferenceFloatList returns a bound list of float64 values that is managed by the application preferences. Changes to this value will be saved to application storage and when the app starts the previous values will be read.


<div class="since">Since: <code>
2.6</code></div>
