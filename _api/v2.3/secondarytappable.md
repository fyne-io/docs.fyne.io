---
layout: page
tags: [api]
title: Fyne API "fyne.SecondaryTappable"
---

# fyne.SecondaryTappable
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type SecondaryTappable

```go
type SecondaryTappable interface {
	TappedSecondary(*PointEvent)
}
```

SecondaryTappable describes a CanvasObject that can be right-clicked or long-tapped.
