---
layout: page
tags: [api]
title: Fyne API "binding.Bool"
---

# binding.Bool
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Bool

```go
type Bool interface {
	DataItem
	Get() (bool, error)
	Set(bool) error
}
```

Bool supports binding a bool value.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindPreferenceBool

```go
func BindPreferenceBool(key string, p fyne.Preferences) Bool
```
BindPreferenceBool returns a bindable bool value that is managed by the application preferences. Changes to this value will be saved to application storage and when the app starts the previous values will be read.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewBool

```go
func NewBool() Bool
```
NewBool returns a bindable bool value that is managed internally.


<div class="since">Since: <code>
2.0</code></div>

#### func  StringToBool

```go
func StringToBool(str String) Bool
```
StringToBool creates a binding that connects a String data item to a Bool. Changes to the String will be parsed and pushed to the Bool if the parse was successful, and setting the Bool update the String binding.


<div class="since">Since: <code>
2.0</code></div>

#### func  StringToBoolWithFormat

```go
func StringToBoolWithFormat(str String, format string) Bool
```
StringToBoolWithFormat creates a binding that connects a String data item to a Bool and is presented using the specified format. Changes to the Bool will be parsed and if the format matches and the parse is successful it will be pushed to the String. Setting the Bool will push a formatted value into the String.


<div class="since">Since: <code>
2.0</code></div>
