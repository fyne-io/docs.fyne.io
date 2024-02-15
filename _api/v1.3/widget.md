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

#### type Widget

```go
type Widget interface {
	CanvasObject

	CreateRenderer() WidgetRenderer
}
```

Widget defines the standard behaviours of any widget. This extends the CanvasObject - a widget behaves in the same basic way but will encapsulate many child objects to create the rendered widget.
