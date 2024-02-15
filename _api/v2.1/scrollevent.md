---
layout: page
tags: [api]
title: Fyne API "fyne.ScrollEvent"
---

# fyne.ScrollEvent
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ScrollEvent

```go
type ScrollEvent struct {
	PointEvent
	Scrolled Delta
}
```

ScrollEvent defines the parameters of a pointer or other scroll event. The DeltaX and DeltaY represent how large the scroll was in two dimensions.
