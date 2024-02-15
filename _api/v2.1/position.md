---
layout: page
tags: [api]
title: Fyne API "fyne.Position"
---

# fyne.Position
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Position

```go
type Position struct {
	X float32 // The position from the parent's left edge
	Y float32 // The position from the parent's top edge
}
```

Position describes a generic X, Y coordinate relative to a parent Canvas or CanvasObject.

#### func  NewPos

```go
func NewPos(x float32, y float32) Position
```
NewPos returns a newly allocated Position representing the specified coordinates.

#### func (Position) Add

```go
func (p Position) Add(v Vector2) Position
```
Add returns a new Position that is the result of offsetting the current position by p2 X and Y.

#### func (Position) Components

```go
func (p Position) Components() (float32, float32)
```
Components returns the X and Y elements of this Position

#### func (Position) IsZero

```go
func (p Position) IsZero() bool
```
IsZero returns whether the Position is at the zero-point.

#### func (Position) Subtract

```go
func (p Position) Subtract(v Vector2) Position
```
Subtract returns a new Position that is the result of offsetting the current position by p2 -X and -Y.
