---
layout: page
tags: [api]
title: Fyne API "fyne.Delta"
package: fyne.io/fyne/v2
---

# fyne.Delta
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Delta

```go
type Delta struct {
	DX, DY float32
}
```

Delta is a generic X, Y coordinate, size or movement representation.

#### func  NewDelta

```go
func NewDelta(dx float32, dy float32) Delta
```
NewDelta returns a newly allocated Delta representing a movement in the X and Y axis.

#### func (Delta) Components

```go
func (v Delta) Components() (float32, float32)
```
Components returns the X and Y elements of this Delta.

#### func (Delta) IsZero

```go
func (v Delta) IsZero() bool
```
IsZero returns whether the Position is at the zero-point.
