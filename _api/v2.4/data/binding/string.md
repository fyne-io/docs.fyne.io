---
layout: page
tags: [api]
title: Fyne API "binding.String"
package: fyne.io/fyne/v2/data/binding
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

#### func  NewSprintf

```go
func NewSprintf(format string, b ...DataItem) String
```
NewSprintf returns a String binding that format its content using the format string and the provide additional parameter that must be other data bindings. This data binding use fmt.Sprintf and fmt.Scanf internally and will have all the same limitation as those function.


<div class="since">Since: <code>
2.2</code></div>

#### func  NewString

```go
func NewString() String
```
NewString returns a bindable string value that is managed internally.


<div class="since">Since: <code>
2.0</code></div>

#### func  StringToStringWithFormat

```go
func StringToStringWithFormat(str String, format string) String
```
StringToStringWithFormat creates a binding that converts a string to another string using the specified format. Changes to the returned String will be pushed to the passed in String and setting a new string value will parse and set the underlying String if it matches the format and the parse was successful.


<div class="since">Since: <code>
2.2</code></div>

#### func  URIToString

```go
func URIToString(v URI) String
```
URIToString creates a binding that connects a URI data item to a String. Changes to the URI will be pushed to the String and setting the string will parse and set the URI if the parse was successful.


<div class="since">Since: <code>
2.1</code></div>
