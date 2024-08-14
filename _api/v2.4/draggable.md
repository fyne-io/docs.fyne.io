---
layout: page
tags: [api]
title: Fyne API "fyne.Draggable"
package: fyne.io/fyne/v2
---

# fyne.Draggable
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Draggable

```go
type Draggable interface {
	Dragged(*DragEvent)
	DragEnd()
}
```

Draggable indicates that a CanvasObject can be dragged. This is used for any item that the user has indicated should be moved across the screen.
