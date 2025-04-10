---
layout: page
tags: [api]
title: Fyne API "binding.IntList"
package: fyne.io/fyne/v2/data/binding
---

# binding.IntList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type IntList

```go
type IntList interface {
	DataList

	Append(value int) error
	Get() ([]int, error)
	GetValue(index int) (int, error)
	Prepend(value int) error
	Remove(value int) error
	Set(list []int) error
	SetValue(index int, value int) error
}
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

#### func  NewIntList

```go
func NewIntList() IntList
```
NewIntList returns a bindable list of int values.


<div class="since">Since: <code>
2.0</code></div>
