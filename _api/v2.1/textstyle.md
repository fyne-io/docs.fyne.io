---
layout: page
tags: [api]
title: Fyne API "fyne.TextStyle"
---

# fyne.TextStyle
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type TextStyle

```go
type TextStyle struct {
	Bold      bool // Should text be bold
	Italic    bool // Should text be italic
	Monospace bool // Use the system monospace font instead of regular

	// Since: 2.1
	TabWidth int // Width of tabs in spaces
}
```

TextStyle represents the styles that can be applied to a text canvas object or text based widget.
