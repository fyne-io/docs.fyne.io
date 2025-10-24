---
tags: [api]
title: fyne.TextStyle
slug: textstyle

aliases:
- /api/v2.7//textstyle

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type TextStyle struct {
	Bold      bool // Should text be bold
	Italic    bool // Should text be italic
	Monospace bool // Use the system monospace font instead of regular
	// Since: 2.2
	Symbol bool // Use the system symbol font.
	// Since: 2.1
	TabWidth int // Width of tabs in spaces
	// Since: 2.5
	// Currently only supported by [fyne.io/fyne/v2/widget.TextGrid].
	Underline bool // Should text be underlined.
}
```

TextStyle represents the styles that can be applied to a text canvas object or text based widget.
