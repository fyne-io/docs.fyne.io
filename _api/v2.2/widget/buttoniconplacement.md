---
layout: page
tags: [api]
title: Fyne API "widget.ButtonIconPlacement"
---

# widget.ButtonIconPlacement
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
