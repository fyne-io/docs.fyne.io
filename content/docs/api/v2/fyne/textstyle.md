---
tags: [api]
title: fyne.TextStyle
slug: textstyle

aliases:
- /api/v2/fyne/textstyle.html
- /api/v2.0//textstyle
- /api/v2.0//textstyle.html
- /api/v2.1//textstyle
- /api/v2.1//textstyle.html
- /api/v2.2//textstyle
- /api/v2.2//textstyle.html
- /api/v2.3//textstyle
- /api/v2.3//textstyle.html
- /api/v2.4//textstyle
- /api/v2.4//textstyle.html
- /api/v2.5//textstyle
- /api/v2.5//textstyle.html
- /api/v2.6//textstyle
- /api/v2.6//textstyle.html
- /api/v2.7//textstyle
- /api/v2.7//textstyle.html

package: fyne.io/fyne/v2
---


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
