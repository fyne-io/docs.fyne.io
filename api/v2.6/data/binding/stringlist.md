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
type StringList interface {
	DataList

	Append(value string) error
	Get() ([]string, error)
	GetValue(index int) (string, error)
	Prepend(value string) error
	Remove(value string) error
	Set(list []string) error
	SetValue(index int, value string) error
}
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

#### func  NewStringList

```go
func NewStringList() StringList
```
NewStringList returns a bindable list of string values.


<div class="since">Since: <code>
2.0</code></div>
