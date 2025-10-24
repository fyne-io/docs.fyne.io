---
tags: [api]
title: fyne.Layout
slug: layout

aliases:
- /api/v2.0//layout
- /api/v2.1//layout
- /api/v2.2//layout
- /api/v2.3//layout
- /api/v2.4//layout
- /api/v2.5//layout
- /api/v2.6//layout
- /api/v2.7//layout

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Layout

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
