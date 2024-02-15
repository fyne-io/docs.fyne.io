---
layout: page
tags: [api]
title: Fyne API "fyne.Layout"
---

# fyne.Layout
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Layout

```go
type Layout interface {
	// Layout will manipulate the listed CanvasObjects Size and Position
	// to fit within the specified size.
	Layout([]CanvasObject, Size)
	// MinSize calculates the smallest size that will fit the listed
	// CanvasObjects using this Layout algorithm.
	MinSize(objects []CanvasObject) Size
}
```

Layout defines how CanvasObjects may be laid out in a specified Size.
