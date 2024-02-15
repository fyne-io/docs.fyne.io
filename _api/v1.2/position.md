---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Position

```go
type Position struct {
	X int // The position from the parent's left edge
	Y int // The position from the parent's top edge
}
```

Position describes a generic X, Y coordinate relative to a parent Canvas or CanvasObject.

#### func  NewPos

```go
func NewPos(x int, y int) Position
```
NewPos returns a newly allocated Position representing the specified coordinates.

#### func (Position) Add

```go
func (p1 Position) Add(p2 Position) Position
```
Add returns a new Position that is the result of offsetting the current position by p2 X and Y.

#### func (Position) Subtract

```go
func (p1 Position) Subtract(p2 Position) Position
```
Subtract returns a new Position that is the result of offsetting the current position by p2 -X and -Y.
