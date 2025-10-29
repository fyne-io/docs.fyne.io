---
tags: [api]
title: binding.Tree
slug: tree

aliases:
- /api/data/binding/tree
- /api/data/binding/tree.html
- /api/v2.0/data/binding/tree
- /api/v2.0/data/binding/tree.html
- /api/v2.1/data/binding/tree
- /api/v2.1/data/binding/tree.html
- /api/v2.2/data/binding/tree
- /api/v2.2/data/binding/tree.html
- /api/v2.3/data/binding/tree
- /api/v2.3/data/binding/tree.html
- /api/v2.4/data/binding/tree
- /api/v2.4/data/binding/tree.html
- /api/v2.5/data/binding/tree
- /api/v2.5/data/binding/tree.html
- /api/v2.6/data/binding/tree
- /api/v2.6/data/binding/tree.html
- /api/v2.7/data/binding/tree
- /api/v2.7/data/binding/tree.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type Tree

```go
type Tree[T any] interface {
	DataTree

	Append(parent, id string, value T) error
	Get() (map[string][]string, map[string]T, error)
	GetValue(id string) (T, error)
	Prepend(parent, id string, value T) error
	Remove(id string) error
	Set(ids map[string][]string, values map[string]T) error
	SetValue(id string, value T) error
}
```

Tree supports binding a tree of values with type T.


<div class="since">Since: <code>
2.7</code></div>

#### func  NewBoolTree

```go
func NewBoolTree() Tree[bool]
```
NewBoolTree returns a bindable tree of bool values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewBytesTree

```go
func NewBytesTree() Tree[[]byte]
```
NewBytesTree returns a bindable tree of []byte values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewFloatTree

```go
func NewFloatTree() Tree[float64]
```
NewFloatTree returns a bindable tree of float64 values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewIntTree

```go
func NewIntTree() Tree[int]
```
NewIntTree returns a bindable tree of int values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewRuneTree

```go
func NewRuneTree() Tree[rune]
```
NewRuneTree returns a bindable tree of rune values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewStringTree

```go
func NewStringTree() Tree[string]
```
NewStringTree returns a bindable tree of string values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewTree

```go
func NewTree[T any](comparator func(T, T) bool) Tree[T]
```
NewTree returns a bindable tree of values with type T.


<div class="since">Since: <code>
2.7</code></div>

#### func  NewURITree

```go
func NewURITree() Tree[fyne.URI]
```
NewURITree returns a bindable tree of fyne.URI values.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewUntypedTree

```go
func NewUntypedTree() Tree[any]
```
NewUntypedTree returns a bindable tree of any values.


<div class="since">Since: <code>
2.5</code></div>
