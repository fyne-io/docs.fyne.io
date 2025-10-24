---
tags: [api]
title: fyne.OverlayStack
slug: overlaystack

aliases:
- /api/v2.7//overlaystack

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

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

OverlayStack is a stack of [CanvasObject]s intended to be used as overlays of a [Canvas].
