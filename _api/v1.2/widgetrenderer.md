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

#### type WidgetRenderer

```go
type WidgetRenderer interface {
	Layout(Size)
	MinSize() Size

	Refresh()
	BackgroundColor() color.Color
	Objects() []CanvasObject
	Destroy()
}
```

WidgetRenderer defines the behaviour of a widget's implementation. This is returned from a widget's declarative object through the Render() function and should be exactly one instance per widget in memory.
