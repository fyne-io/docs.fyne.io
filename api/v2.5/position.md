---
layout: page
tags: [api]
title: Fyne API "fyne.Position"
package: fyne.io/fyne/v2
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

#### func  NewSquareOffsetPos

```go
func NewSquareOffsetPos(length float32) Position
```
NewSquareOffsetPos returns a newly allocated Position with the same x and y position.


<div class="since">Since: <code>
2.4</code></div>

#### func (Position) Add

```go
func (p Position) Add(v Vector2) Position
```
Add returns a new Position that is the result of offsetting the current position by p2 X and Y.

#### func (Position) AddXY

```go
func (p Position) AddXY(x, y float32) Position
```
AddXY returns a new Position by adding x and y to the current one.

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

#### func (Position) SubtractXY

```go
func (p Position) SubtractXY(x, y float32) Position
```
SubtractXY returns a new Position by subtracting x and y from the current one.
