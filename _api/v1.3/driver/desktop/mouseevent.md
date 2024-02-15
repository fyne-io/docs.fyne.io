---
layout: page
tags: [api]
title: Fyne API desktop
---

# desktop
---
```go
import "fyne.io/fyne/driver/desktop"
```

## Usage

#### type MouseEvent

```go
type MouseEvent struct {
	fyne.PointEvent
	Button   MouseButton
	Modifier Modifier
}
```

MouseEvent contains data relating to desktop mouse events
