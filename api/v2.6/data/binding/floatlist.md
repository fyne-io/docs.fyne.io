---
layout: page
tags: [api]
title: Fyne API "binding.FloatList"
package: fyne.io/fyne/v2/data/binding
---

# binding.FloatList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type FloatList

```go
type FloatList interface {
	DataList

	Append(value float64) error
	Get() ([]float64, error)
	GetValue(index int) (float64, error)
	Prepend(value float64) error
	Remove(value float64) error
	Set(list []float64) error
	SetValue(index int, value float64) error
}
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

#### func  NewFloatList

```go
func NewFloatList() FloatList
```
NewFloatList returns a bindable list of float64 values.


<div class="since">Since: <code>
2.0</code></div>
