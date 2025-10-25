---
tags: [api]
title: widget.ButtonIconPlacement
slug: buttoniconplacement

aliases:
- /api/v2/widget/buttoniconplacement.html
- /api/v2.0/widget/buttoniconplacement
- /api/v2.0/widget/buttoniconplacement.html
- /api/v2.1/widget/buttoniconplacement
- /api/v2.1/widget/buttoniconplacement.html
- /api/v2.2/widget/buttoniconplacement
- /api/v2.2/widget/buttoniconplacement.html
- /api/v2.3/widget/buttoniconplacement
- /api/v2.3/widget/buttoniconplacement.html
- /api/v2.4/widget/buttoniconplacement
- /api/v2.4/widget/buttoniconplacement.html
- /api/v2.5/widget/buttoniconplacement
- /api/v2.5/widget/buttoniconplacement.html
- /api/v2.6/widget/buttoniconplacement
- /api/v2.6/widget/buttoniconplacement.html
- /api/v2.7/widget/buttoniconplacement
- /api/v2.7/widget/buttoniconplacement.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ButtonIconPlacement

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
