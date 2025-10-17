---
layout: page
tags: [api]
title: Fyne API "binding.StringList"
package: fyne.io/fyne/v2/data/binding
---

# binding.StringList
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
