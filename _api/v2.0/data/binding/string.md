---
layout: page
tags: [api]
title: Fyne API "binding.String"
---

# binding.String
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type String

```go
type String interface {
	DataItem
	Get() (string, error)
	Set(string) error
}
```

String supports binding a string value.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindPreferenceString

```go
func BindPreferenceString(key string, p fyne.Preferences) String
```
BindPreferenceString returns a bindable string value that is managed by the application preferences. Changes to this value will be saved to application storage and when the app starts the previous values will be read.


<div class="since">Since: <code>
2.0</code></div>

#### func  BoolToString

```go
func BoolToString(v Bool) String
```
BoolToString creates a binding that connects a Bool data item to a String. Changes to the Bool will be pushed to the String and setting the string will parse and set the Bool if the parse was successful.


<div class="since">Since: <code>
2.0</code></div>

#### func  BoolToStringWithFormat

```go
func BoolToStringWithFormat(v Bool, format string) String
```
BoolToStringWithFormat creates a binding that connects a Bool data item to a String and is presented using the specified format. Changes to the Bool will be pushed to the String and setting the string will parse and set the Bool if the string matches the format and its parse was successful.


<div class="since">Since: <code>
2.0</code></div>

#### func  FloatToString

```go
func FloatToString(v Float) String
```
FloatToString creates a binding that connects a Float data item to a String. Changes to the Float will be pushed to the String and setting the string will parse and set the Float if the parse was successful.


<div class="since">Since: <code>
2.0</code></div>

#### func  FloatToStringWithFormat

```go
func FloatToStringWithFormat(v Float, format string) String
```
FloatToStringWithFormat creates a binding that connects a Float data item to a String and is presented using the specified format. Changes to the Float will be pushed to the String and setting the string will parse and set the Float if the string matches the format and its parse was successful.


<div class="since">Since: <code>
2.0</code></div>

#### func  IntToString

```go
func IntToString(v Int) String
```
IntToString creates a binding that connects a Int data item to a String. Changes to the Int will be pushed to the String and setting the string will parse and set the Int if the parse was successful.


<div class="since">Since: <code>
2.0</code></div>

#### func  IntToStringWithFormat

```go
func IntToStringWithFormat(v Int, format string) String
```
IntToStringWithFormat creates a binding that connects a Int data item to a String and is presented using the specified format. Changes to the Int will be pushed to the String and setting the string will parse and set the Int if the string matches the format and its parse was successful.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewString

```go
func NewString() String
```
NewString returns a bindable string value that is managed internally.


<div class="since">Since: <code>
2.0</code></div>
