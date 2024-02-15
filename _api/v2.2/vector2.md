---
layout: page
tags: [api]
title: Fyne API "fyne.Vector2"
---

# fyne.Vector2
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
