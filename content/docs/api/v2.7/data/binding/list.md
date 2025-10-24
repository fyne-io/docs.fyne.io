---
tags: [api]
title: binding.List
slug: list

aliases:
- /api/v2.7/data/binding/list

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type List[T any] interface {
	DataList

	Append(value T) error
	Get() ([]T, error)
	GetValue(index int) (T, error)
	Prepend(value T) error
	Remove(value T) error
	Set(list []T) error
	SetValue(index int, value T) error
}
```

List supports binding a list of values with type T.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func NewBoolList() List[bool]
```
NewBoolList returns a bindable list of bool values.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func NewBytesList() List[[]byte]
```
NewBytesList returns a bindable list of []byte values.


<div class="since">Since: <code>
2.2</code></div>

###

```go
func NewFloatList() List[float64]
```
NewFloatList returns a bindable list of float64 values.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func NewIntList() List[int]
```
NewIntList returns a bindable list of int values.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func NewList[T any](comparator func(T, T) bool) List[T]
```
NewList returns a bindable list of values with type T.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func NewRuneList() List[rune]
```
NewRuneList returns a bindable list of rune values.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func NewStringList() List[string]
```
NewStringList returns a bindable list of string values.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func NewURIList() List[fyne.URI]
```
NewURIList returns a bindable list of fyne.URI values.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func NewUntypedList() List[any]
```
NewUntypedList returns a bindable list of any values.


<div class="since">Since: <code>
2.1</code></div>
