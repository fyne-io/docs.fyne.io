---
tags: [api]
title: binding.StringList
slug: stringlist

aliases:
- /api/v2.0/data/binding/stringlist
- /api/v2.1/data/binding/stringlist
- /api/v2.2/data/binding/stringlist
- /api/v2.3/data/binding/stringlist
- /api/v2.4/data/binding/stringlist
- /api/v2.5/data/binding/stringlist
- /api/v2.6/data/binding/stringlist
- /api/v2.7/data/binding/stringlist

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type StringList

```go
type StringList = List[string]
```

StringList supports binding a list of string values.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindPreferenceStringList

```go
func BindPreferenceStringList(key string, p fyne.Preferences) StringList
```
BindPreferenceStringList returns a bound list of string values that is managed by the application preferences. Changes to this value will be saved to application storage and when the app starts the previous values will be read.


<div class="since">Since: <code>
2.6</code></div>
