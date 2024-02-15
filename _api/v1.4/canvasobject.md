---
layout: page
tags: [api]
title: Fyne API "fyne.CanvasObject"
---

# fyne.CanvasObject
---
```go
import "fyne.io/fyne"
```

## Usage

#### type CanvasObject

```go
type CanvasObject interface {

	// MinSize returns the minimum size this object needs to be drawn.
	MinSize() Size
	// Move moves this object to the given position relative to its parent.
	// This should only be called if your object is not in a container with a layout manager.
	Move(Position)
	// Position returns the current position of the object relative to its parent.
	Position() Position
	// Resize resizes this object to the given size.
	// This should only be called if your object is not in a container with a layout manager.
	Resize(Size)
	// Size returns the current size of this object.
	Size() Size

	// Hide hides this object.
	Hide()
	// Visible returns whether this object is visible or not.
	Visible() bool
	// Show shows this object.
	Show()

	// Refresh must be called if this object should be redrawn because its inner state changed.
	Refresh()
}
```

CanvasObject describes any graphical object that can be added to a canvas. Objects have a size and position that can be controlled through this API. MinSize is used to determine the minimum size which this object should be displayed. An object will be visible by default but can be hidden with Hide() and re-shown with Show().

Note: If this object is controlled as part of a Layout you should not call Resize(Size) or Move(Position).
