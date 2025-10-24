---
tags: [api]
title: widget.Orientation
slug: orientation

aliases:
- /api/v2.0/widget/orientation
- /api/v2.1/widget/orientation
- /api/v2.2/widget/orientation
- /api/v2.3/widget/orientation
- /api/v2.4/widget/orientation
- /api/v2.5/widget/orientation
- /api/v2.6/widget/orientation
- /api/v2.7/widget/orientation

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Orientation

```go
type Orientation int
```

Orientation controls the horizontal/vertical layout of a widget

```go
const (
	Horizontal Orientation = 0
	Vertical   Orientation = 1

	// Adaptive will switch between horizontal and vertical layouts according to device orientation.
	// This orientation is not always supported and interpretation can vary per-widget.
	//
	// Since: 2.5
	Adaptive Orientation = 2
)
```
Orientation constants to control widget layout
