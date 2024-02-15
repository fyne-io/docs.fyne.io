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

#### type CanvasObject

```go
type CanvasObject interface {
	// geometry
	Size() Size
	Resize(Size)
	Position() Position
	Move(Position)
	MinSize() Size

	// visibility
	Visible() bool
	// Show shows this object.
	Show()
	// Hide hides this object.
	Hide()

	Refresh()
}
```

CanvasObject describes any graphical object that can be added to a canvas. Objects have a size and position that can be controlled through this API. MinSize is used to determine the minimum size which this object should be displayed. An object will be visible by default but can be hidden with Hide() and re-shown with Show().

Note: If this object is controlled as part of a Layout you should not call Resize(Size) or Move(Position).
