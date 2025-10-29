---
tags: [api]
title: fyne.Vector2
slug: vector2

aliases:
- /api//vector2
- /api//vector2.html
- /api/v2.0//vector2
- /api/v2.0//vector2.html
- /api/v2.1//vector2
- /api/v2.1//vector2.html
- /api/v2.2//vector2
- /api/v2.2//vector2.html
- /api/v2.3//vector2
- /api/v2.3//vector2.html
- /api/v2.4//vector2
- /api/v2.4//vector2.html
- /api/v2.5//vector2
- /api/v2.5//vector2.html
- /api/v2.6//vector2
- /api/v2.6//vector2.html
- /api/v2.7//vector2
- /api/v2.7//vector2.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Vector2

```go
type Vector2 interface {
	Components() (float32, float32)
	IsZero() bool
}
```

Vector2 marks geometry types that can operate as a coordinate vector.
