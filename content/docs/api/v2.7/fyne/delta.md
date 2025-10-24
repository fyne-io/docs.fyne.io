---
tags: [api]
title: fyne.Delta
slug: delta

aliases:
- /api/v2.7//delta

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Delta struct {
	DX, DY float32
}
```

Delta is a generic X, Y coordinate, size or movement representation.

###

```go
func NewDelta(dx float32, dy float32) Delta
```
NewDelta returns a newly allocated [Delta] representing a movement in the X and Y axis.

###

```go
func (v Delta) Components() (float32, float32)
```
Components returns the X and Y elements of v.

###

```go
func (v Delta) IsZero() bool
```
IsZero returns whether the Position is at the zero-point.
