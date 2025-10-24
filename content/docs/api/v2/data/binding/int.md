---
tags: [api]
title: binding.Int
slug: int

aliases:
- /api/v2.0/data/binding/int
- /api/v2.1/data/binding/int
- /api/v2.2/data/binding/int
- /api/v2.3/data/binding/int
- /api/v2.4/data/binding/int
- /api/v2.5/data/binding/int
- /api/v2.6/data/binding/int
- /api/v2.7/data/binding/int

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Int

```go
type Int = Item[int]
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

#### func  FloatToInt

```go
func FloatToInt(v Float) Int
```
FloatToInt creates a binding that connects a Float data item to an Int.


<div class="since">Since: <code>
2.5</code></div>

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
