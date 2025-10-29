---
tags: [api]
title: fyne.OverlayStack
slug: overlaystack

aliases:
- /api//overlaystack
- /api//overlaystack.html
- /api/v2.0//overlaystack
- /api/v2.0//overlaystack.html
- /api/v2.1//overlaystack
- /api/v2.1//overlaystack.html
- /api/v2.2//overlaystack
- /api/v2.2//overlaystack.html
- /api/v2.3//overlaystack
- /api/v2.3//overlaystack.html
- /api/v2.4//overlaystack
- /api/v2.4//overlaystack.html
- /api/v2.5//overlaystack
- /api/v2.5//overlaystack.html
- /api/v2.6//overlaystack
- /api/v2.6//overlaystack.html
- /api/v2.7//overlaystack
- /api/v2.7//overlaystack.html

package: fyne.io/fyne/v2
---


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

OverlayStack is a stack of [CanvasObject]s intended to be used as overlays of a [Canvas].
