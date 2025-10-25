---
tags: [api]
title: binding.FloatList
slug: floatlist

aliases:
- /api/v2/data/binding/floatlist.html
- /api/v2.0/data/binding/floatlist
- /api/v2.0/data/binding/floatlist.html
- /api/v2.1/data/binding/floatlist
- /api/v2.1/data/binding/floatlist.html
- /api/v2.2/data/binding/floatlist
- /api/v2.2/data/binding/floatlist.html
- /api/v2.3/data/binding/floatlist
- /api/v2.3/data/binding/floatlist.html
- /api/v2.4/data/binding/floatlist
- /api/v2.4/data/binding/floatlist.html
- /api/v2.5/data/binding/floatlist
- /api/v2.5/data/binding/floatlist.html
- /api/v2.6/data/binding/floatlist
- /api/v2.6/data/binding/floatlist.html
- /api/v2.7/data/binding/floatlist
- /api/v2.7/data/binding/floatlist.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type FloatList

```go
type FloatList = List[float64]
```

FloatList supports binding a list of float64 values.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindPreferenceFloatList

```go
func BindPreferenceFloatList(key string, p fyne.Preferences) FloatList
```
BindPreferenceFloatList returns a bound list of float64 values that is managed by the application preferences. Changes to this value will be saved to application storage and when the app starts the previous values will be read.


<div class="since">Since: <code>
2.6</code></div>
