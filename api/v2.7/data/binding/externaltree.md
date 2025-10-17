---
layout: page
tags: [api]
title: Fyne API "binding.ExternalTree"
package: fyne.io/fyne/v2/data/binding
---

# binding.ExternalTree
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalTree

```go
type ExternalTree[T any] interface {
	Tree[T]

	Reload() error
}
```

ExternalTree supports binding a tree of values, of type T, from an external variable.


<div class="since">Since: <code>
2.7</code></div>

#### func  BindBoolTree

```go
func BindBoolTree(ids *map[string][]string, v *map[string]bool) ExternalTree[bool]
```
BindBoolTree returns a bound tree of bool values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindBytesTree

```go
func BindBytesTree(ids *map[string][]string, v *map[string][]byte) ExternalTree[[]byte]
```
BindBytesTree returns a bound tree of []byte values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindFloatTree

```go
func BindFloatTree(ids *map[string][]string, v *map[string]float64) ExternalTree[float64]
```
BindFloatTree returns a bound tree of float64 values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindIntTree

```go
func BindIntTree(ids *map[string][]string, v *map[string]int) ExternalTree[int]
```
BindIntTree returns a bound tree of int values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindRuneTree

```go
func BindRuneTree(ids *map[string][]string, v *map[string]rune) ExternalTree[rune]
```
BindRuneTree returns a bound tree of rune values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindStringTree

```go
func BindStringTree(ids *map[string][]string, v *map[string]string) ExternalTree[string]
```
BindStringTree returns a bound tree of string values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindTree

```go
func BindTree[T any](ids *map[string][]string, v *map[string]T, comparator func(T, T) bool) ExternalTree[T]
```
BindTree returns a bound tree of values with type T, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.7</code></div>

#### func  BindURITree

```go
func BindURITree(ids *map[string][]string, v *map[string]fyne.URI) ExternalTree[fyne.URI]
```
BindURITree returns a bound tree of fyne.URI values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>

#### func  BindUntypedTree

```go
func BindUntypedTree(ids *map[string][]string, v *map[string]any) ExternalTree[any]
```
BindUntypedTree returns a bound tree of any values, based on the contents of the passed values. The ids map specifies how each item relates to its parent (with id ""), with the values being in the v map. If your code changes the content of the maps this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.4</code></div>
