---
layout: page
tags: [api]
title: Fyne API "fyne.WidgetRenderer"
---

# fyne.WidgetRenderer
---
```go
import "fyne.io/fyne"
```

## Usage

#### type WidgetRenderer

```go
type WidgetRenderer interface {
	// BackgroundColor returns the color that should be used to draw the background of this renderer's widget.
	//
	// Deprecated: Widgets will no longer have a background to support hover and selection indication in collection widgets.
	// If a widget requires a background color or image, this can be achieved by using a canvas.Rect or canvas.Image
	// as the first child of a MaxLayout, followed by the rest of the widget components.
	BackgroundColor() color.Color
	// Destroy is for internal use.
	Destroy()
	// Layout is a hook that is called if the widget needs to be laid out.
	// This should never call Refresh.
	Layout(Size)
	// MinSize returns the minimum size of the widget that is rendered by this renderer.
	MinSize() Size
	// Objects returns all objects that should be drawn.
	Objects() []CanvasObject
	// Refresh is a hook that is called if the widget has updated and needs to be redrawn.
	// This might trigger a Layout.
	Refresh()
}
```

WidgetRenderer defines the behaviour of a widget's implementation. This is returned from a widget's declarative object through the CreateRenderer() function and should be exactly one instance per widget in memory.
