---
tags: [api]
title: fyne.Widget
slug: widget

aliases:
- /api/v2.0//widget
- /api/v2.1//widget
- /api/v2.2//widget
- /api/v2.3//widget
- /api/v2.4//widget
- /api/v2.5//widget
- /api/v2.6//widget
- /api/v2.7//widget

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Widget

```go
type Widget interface {
	CanvasObject

	// CreateRenderer returns a new [WidgetRenderer] for this widget.
	// This should not be called by regular code, it is used internally to render a widget.
	CreateRenderer() WidgetRenderer
}
```

Widget defines the standard behaviours of any widget. This extends [CanvasObject]. A widget behaves in the same basic way but will encapsulate many child objects to create the rendered widget.
