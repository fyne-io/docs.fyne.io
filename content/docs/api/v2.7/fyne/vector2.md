---
tags: [api]
title: fyne.Vector2
slug: vector2

aliases:
- /api/v2.7//vector2

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Vector2 interface {
	Components() (float32, float32)
	IsZero() bool
}
```

Vector2 marks geometry types that can operate as a coordinate vector.
