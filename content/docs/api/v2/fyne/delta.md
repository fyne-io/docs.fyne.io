---
tags: [api]
title: fyne.Delta
slug: delta

aliases:
- /api/v2.0//delta
- /api/v2.1//delta
- /api/v2.2//delta
- /api/v2.3//delta
- /api/v2.4//delta
- /api/v2.5//delta
- /api/v2.6//delta
- /api/v2.7//delta

package: fyne.io/fyne/v2
---


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
NewDelta returns a newly allocated [Delta] representing a movement in the X and Y axis.

#### func (Delta) Components

```go
func (v Delta) Components() (float32, float32)
```
Components returns the X and Y elements of v.

#### func (Delta) IsZero

```go
func (v Delta) IsZero() bool
```
IsZero returns whether the Position is at the zero-point.
