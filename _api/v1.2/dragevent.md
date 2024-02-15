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

#### type DragEvent

```go
type DragEvent struct {
	PointEvent
	DraggedX, DraggedY int
}
```

DragEvent defines the parameters of a pointer or other drag event. The DraggedX and DraggedY fields show how far the item was dragged since the last event.
