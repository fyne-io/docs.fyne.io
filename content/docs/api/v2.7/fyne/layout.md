---
tags: [api]
title: fyne.Layout
slug: layout

aliases:
- /api/v2.7//layout

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Layout interface {
	// Layout will manipulate the listed [CanvasObject]s Size and Position
	// to fit within the specified size.
	Layout([]CanvasObject, Size)
	// MinSize calculates the smallest size that will fit the listed
	// [CanvasObject]s using this Layout algorithm.
	MinSize(objects []CanvasObject) Size
}
```

Layout defines how [CanvasObject]s may be laid out in a specified Size.
