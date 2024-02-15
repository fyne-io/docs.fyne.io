---
layout: page
tags: [api]
title: Fyne API "fyne.Scrollable"
---

# fyne.Scrollable
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Scrollable

```go
type Scrollable interface {
	Scrolled(*ScrollEvent)
}
```

Scrollable describes any CanvasObject that can also be scrolled. This is mostly used to implement the widget.ScrollContainer.
