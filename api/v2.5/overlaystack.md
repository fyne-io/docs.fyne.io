---
layout: page
tags: [api]
title: Fyne API "fyne.OverlayStack"
package: fyne.io/fyne/v2
---

# fyne.OverlayStack
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type OverlayStack

```go
type OverlayStack interface {
	// Add adds an overlay on the top of the overlay stack.
	Add(overlay CanvasObject)
	// List returns the overlays currently on the overlay stack.
	List() []CanvasObject
	// Remove removes the given object and all objects above it from the overlay stack.
	Remove(overlay CanvasObject)
	// Top returns the top-most object of the overlay stack.
	Top() CanvasObject
}
```

OverlayStack is a stack of CanvasObjects intended to be used as overlays of a Canvas.
