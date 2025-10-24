---
tags: [api]
title: widget.ButtonAlign
slug: buttonalign

aliases:
- /api/v2.0/widget/buttonalign
- /api/v2.1/widget/buttonalign
- /api/v2.2/widget/buttonalign
- /api/v2.3/widget/buttonalign
- /api/v2.4/widget/buttonalign
- /api/v2.5/widget/buttonalign
- /api/v2.6/widget/buttonalign
- /api/v2.7/widget/buttonalign

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ButtonAlign

```go
type ButtonAlign int
```

ButtonAlign represents the horizontal alignment of a button.

```go
const (
	// ButtonAlignCenter aligns the icon and the text centrally.
	ButtonAlignCenter ButtonAlign = iota
	// ButtonAlignLeading aligns the icon and the text with the leading edge.
	ButtonAlignLeading
	// ButtonAlignTrailing aligns the icon and the text with the trailing edge.
	ButtonAlignTrailing
)
```
