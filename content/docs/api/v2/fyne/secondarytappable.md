---
tags: [api]
title: fyne.SecondaryTappable
slug: secondarytappable

aliases:
- /api/v2.0//secondarytappable
- /api/v2.1//secondarytappable
- /api/v2.2//secondarytappable
- /api/v2.3//secondarytappable
- /api/v2.4//secondarytappable
- /api/v2.5//secondarytappable
- /api/v2.6//secondarytappable
- /api/v2.7//secondarytappable

package: fyne.io/fyne/v2
---


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

SecondaryTappable describes a [CanvasObject] that can be right-clicked or long-tapped.
