---
tags: [api]
title: widget.ButtonIconPlacement
slug: buttoniconplacement

aliases:
- /api/v2.7/widget/buttoniconplacement

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type ButtonIconPlacement int
```

ButtonIconPlacement represents the ordering of icon & text within a button.

```go
const (
	// ButtonIconLeadingText aligns the icon on the leading edge of the text.
	ButtonIconLeadingText ButtonIconPlacement = iota
	// ButtonIconTrailingText aligns the icon on the trailing edge of the text.
	ButtonIconTrailingText
)
```
