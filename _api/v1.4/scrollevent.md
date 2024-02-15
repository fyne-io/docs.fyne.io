---
layout: page
tags: [api]
title: Fyne API "fyne.ScrollEvent"
---

# fyne.ScrollEvent
---
```go
import "fyne.io/fyne"
```

## Usage

#### type ScrollEvent

```go
type ScrollEvent struct {
	PointEvent
	DeltaX, DeltaY int
}
```

ScrollEvent defines the parameters of a pointer or other scroll event. The DeltaX and DeltaY represent how large the scroll was in two dimensions.
