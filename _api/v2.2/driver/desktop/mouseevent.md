---
layout: page
tags: [api]
title: Fyne API "desktop.MouseEvent"
---

# desktop.MouseEvent
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type MouseEvent

```go
type MouseEvent struct {
	fyne.PointEvent
	Button   MouseButton
	Modifier fyne.KeyModifier
}
```

MouseEvent contains data relating to desktop mouse events
