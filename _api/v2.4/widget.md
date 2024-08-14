---
layout: page
tags: [api]
title: Fyne API "fyne.Widget"
package: fyne.io/fyne/v2
---

# fyne.Widget
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Widget

```go
type Widget interface {
	CanvasObject

	// CreateRenderer returns a new WidgetRenderer for this widget.
	// This should not be called by regular code, it is used internally to render a widget.
	CreateRenderer() WidgetRenderer
}
```

Widget defines the standard behaviours of any widget. This extends the CanvasObject - a widget behaves in the same basic way but will encapsulate many child objects to create the rendered widget.
