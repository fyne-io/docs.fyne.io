---
tags: [api]
title: fyne.SecondaryTappable
slug: secondarytappable

aliases:
- /api/v2.7//secondarytappable

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type SecondaryTappable interface {
	TappedSecondary(*PointEvent)
}
```

SecondaryTappable describes a [CanvasObject] that can be right-clicked or long-tapped.
