---
tags: [api]
title: binding.ExternalList
slug: externallist

aliases:
- /api/v2.7/data/binding/externallist

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type ExternalList[T any] interface {
	List[T]

	Reload() error
}
```

ExternalList supports binding a list of values, with type T, from an external variable.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func BindBoolList(v *[]bool) ExternalList[bool]
```
BindBoolList returns a bound list of bool values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindBytesList(v *[][]byte) ExternalList[[]byte]
```
BindBytesList returns a bound list of []byte values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.2</code></div>

###

```go
func BindFloatList(v *[]float64) ExternalList[float64]
```
BindFloatList returns a bound list of float64 values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindIntList(v *[]int) ExternalList[int]
```
BindIntList returns a bound list of int values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindList[T any](v *[]T, comparator func(T, T) bool) ExternalList[T]
```
BindList returns a bound list of values with type T, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func BindRuneList(v *[]rune) ExternalList[rune]
```
BindRuneList returns a bound list of rune values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindStringList(v *[]string) ExternalList[string]
```
BindStringList returns a bound list of string values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func BindURIList(v *[]fyne.URI) ExternalList[fyne.URI]
```
BindURIList returns a bound list of fyne.URI values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func BindUntypedList(v *[]any) ExternalList[any]
```
BindUntypedList returns a bound list of any values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.1</code></div>
