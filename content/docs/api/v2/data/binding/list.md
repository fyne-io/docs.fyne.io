---
tags: [api]
title: binding.List
slug: list

aliases:
- /api/v2/data/binding/list.html
- /api/v2.0/data/binding/list
- /api/v2.0/data/binding/list.html
- /api/v2.1/data/binding/list
- /api/v2.1/data/binding/list.html
- /api/v2.2/data/binding/list
- /api/v2.2/data/binding/list.html
- /api/v2.3/data/binding/list
- /api/v2.3/data/binding/list.html
- /api/v2.4/data/binding/list
- /api/v2.4/data/binding/list.html
- /api/v2.5/data/binding/list
- /api/v2.5/data/binding/list.html
- /api/v2.6/data/binding/list
- /api/v2.6/data/binding/list.html
- /api/v2.7/data/binding/list
- /api/v2.7/data/binding/list.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type List

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

#### func  NewBoolList

```go
func NewBoolList() List[bool]
```
NewBoolList returns a bindable list of bool values.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewBytesList

```go
func NewBytesList() List[[]byte]
```
NewBytesList returns a bindable list of []byte values.


<div class="since">Since: <code>
2.2</code></div>

#### func  NewFloatList

```go
func NewFloatList() List[float64]
```
NewFloatList returns a bindable list of float64 values.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewIntList

```go
func NewIntList() List[int]
```
NewIntList returns a bindable list of int values.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewList

```go
func NewList[T any](comparator func(T, T) bool) List[T]
```
NewList returns a bindable list of values with type T.


<div class="since">Since: <code>
2.7</code></div>

#### func  NewRuneList

```go
func NewRuneList() List[rune]
```
NewRuneList returns a bindable list of rune values.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewStringList

```go
func NewStringList() List[string]
```
NewStringList returns a bindable list of string values.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewURIList

```go
func NewURIList() List[fyne.URI]
```
NewURIList returns a bindable list of fyne.URI values.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewUntypedList

```go
func NewUntypedList() List[any]
```
NewUntypedList returns a bindable list of any values.


<div class="since">Since: <code>
2.1</code></div>
