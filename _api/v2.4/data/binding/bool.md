---
layout: page
tags: [api]
title: Fyne API "binding.Bool"
package: fyne.io/fyne/v2/data/binding
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

#### func  And

```go
func And(data ...Bool) Bool
```
And returns a Bool binding that return true when all the passed Bool binding are true and false otherwise. It does apply a logical and boolean operation on all passed Bool bindings. This binding is two way. In case of a Set, it will propagate the value identically to all the Bool bindings used for its construction.

Since 2.4

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

#### func  Not

```go
func Not(data Bool) Bool
```
Not returns a Bool binding that invert the value of the given data binding. This is providing the logical Not boolean operation as a data binding.

Since 2.4

#### func  Or

```go
func Or(data ...Bool) Bool
```
Or returns a Bool binding that return true when at least one of the passed Bool binding is true and false otherwise. It does apply a logical or boolean operation on all passed Bool bindings. This binding is two way. In case of a Set, it will propagate the value identically to all the Bool bindings used for its construction.

Since 2.4

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
