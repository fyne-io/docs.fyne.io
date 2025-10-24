---
tags: [api]
title: fyne.WidgetRenderer
slug: widgetrenderer

aliases:
- /api/v2.7//widgetrenderer

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type WidgetRenderer interface {
	// Destroy is a hook that is called when the renderer is being destroyed.
	// This happens at some time after the widget is no longer visible, and
	// once destroyed, a renderer will not be reused.
	// Renderers should dispose and clean up any related resources, if necessary.
	Destroy()
	// Layout is a hook that is called if the widget needs to be laid out.
	// This should never call [Refresh].
	Layout(Size)
	// MinSize returns the minimum size of the widget that is rendered by this renderer.
	MinSize() Size
	// Objects returns all objects that should be drawn.
	Objects() []CanvasObject
	// Refresh is a hook that is called if the widget has updated and needs to be redrawn.
	// This might trigger a [Layout].
	Refresh()
}
```

WidgetRenderer defines the behaviour of a widget's implementation. This is returned from a widget's declarative object through [Widget.CreateRenderer] and should be exactly one instance per widget in memory.
