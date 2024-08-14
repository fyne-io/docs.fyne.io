---
layout: page
tags: [api]
title: Fyne API "binding.Int"
package: fyne.io/fyne/v2/data/binding
---

# binding.Int
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Int

```go
type Int interface {
	DataItem
	Get() (int, error)
	Set(int) error
}
```

Int supports binding a int value.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindPreferenceInt

```go
func BindPreferenceInt(key string, p fyne.Preferences) Int
```
BindPreferenceInt returns a bindable int value that is managed by the application preferences. Changes to this value will be saved to application storage and when the app starts the previous values will be read.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewInt

```go
func NewInt() Int
```
NewInt returns a bindable int value that is managed internally.


<div class="since">Since: <code>
2.0</code></div>

#### func  StringToInt

```go
func StringToInt(str String) Int
```
StringToInt creates a binding that connects a String data item to a Int. Changes to the String will be parsed and pushed to the Int if the parse was successful, and setting the Int update the String binding.


<div class="since">Since: <code>
2.0</code></div>

#### func  StringToIntWithFormat

```go
func StringToIntWithFormat(str String, format string) Int
```
StringToIntWithFormat creates a binding that connects a String data item to a Int and is presented using the specified format. Changes to the Int will be parsed and if the format matches and the parse is successful it will be pushed to the String. Setting the Int will push a formatted value into the String.


<div class="since">Since: <code>
2.0</code></div>
