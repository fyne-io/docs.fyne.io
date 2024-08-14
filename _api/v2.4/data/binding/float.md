---
layout: page
tags: [api]
title: Fyne API "binding.Float"
package: fyne.io/fyne/v2/data/binding
---

# binding.Float
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Float

```go
type Float interface {
	DataItem
	Get() (float64, error)
	Set(float64) error
}
```

Float supports binding a float64 value.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindPreferenceFloat

```go
func BindPreferenceFloat(key string, p fyne.Preferences) Float
```
BindPreferenceFloat returns a bindable float64 value that is managed by the application preferences. Changes to this value will be saved to application storage and when the app starts the previous values will be read.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewFloat

```go
func NewFloat() Float
```
NewFloat returns a bindable float64 value that is managed internally.


<div class="since">Since: <code>
2.0</code></div>

#### func  StringToFloat

```go
func StringToFloat(str String) Float
```
StringToFloat creates a binding that connects a String data item to a Float. Changes to the String will be parsed and pushed to the Float if the parse was successful, and setting the Float update the String binding.


<div class="since">Since: <code>
2.0</code></div>

#### func  StringToFloatWithFormat

```go
func StringToFloatWithFormat(str String, format string) Float
```
StringToFloatWithFormat creates a binding that connects a String data item to a Float and is presented using the specified format. Changes to the Float will be parsed and if the format matches and the parse is successful it will be pushed to the String. Setting the Float will push a formatted value into the String.


<div class="since">Since: <code>
2.0</code></div>
