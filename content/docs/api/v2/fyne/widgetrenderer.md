---
tags: [api]
title: fyne.WidgetRenderer
slug: widgetrenderer

aliases:
- /api//widgetrenderer
- /api//widgetrenderer.html
- /api/v2.0//widgetrenderer
- /api/v2.0//widgetrenderer.html
- /api/v2.1//widgetrenderer
- /api/v2.1//widgetrenderer.html
- /api/v2.2//widgetrenderer
- /api/v2.2//widgetrenderer.html
- /api/v2.3//widgetrenderer
- /api/v2.3//widgetrenderer.html
- /api/v2.4//widgetrenderer
- /api/v2.4//widgetrenderer.html
- /api/v2.5//widgetrenderer
- /api/v2.5//widgetrenderer.html
- /api/v2.6//widgetrenderer
- /api/v2.6//widgetrenderer.html
- /api/v2.7//widgetrenderer
- /api/v2.7//widgetrenderer.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type WidgetRenderer

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
