---
tags: [api]
title: widget.TextGridStyle
slug: textgridstyle

aliases:
- /api/widget/textgridstyle
- /api/widget/textgridstyle.html
- /api/v2.0/widget/textgridstyle
- /api/v2.0/widget/textgridstyle.html
- /api/v2.1/widget/textgridstyle
- /api/v2.1/widget/textgridstyle.html
- /api/v2.2/widget/textgridstyle
- /api/v2.2/widget/textgridstyle.html
- /api/v2.3/widget/textgridstyle
- /api/v2.3/widget/textgridstyle.html
- /api/v2.4/widget/textgridstyle
- /api/v2.4/widget/textgridstyle.html
- /api/v2.5/widget/textgridstyle
- /api/v2.5/widget/textgridstyle.html
- /api/v2.6/widget/textgridstyle
- /api/v2.6/widget/textgridstyle.html
- /api/v2.7/widget/textgridstyle
- /api/v2.7/widget/textgridstyle.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type TextGridStyle

```go
type TextGridStyle interface {
	Style() fyne.TextStyle
	TextColor() color.Color
	BackgroundColor() color.Color
}
```

TextGridStyle defines a style that can be applied to a TextGrid cell.

```go
var (
	// TextGridStyleDefault is a default style for test grid cells
	TextGridStyleDefault TextGridStyle
	// TextGridStyleWhitespace is the style used for whitespace characters, if enabled
	TextGridStyleWhitespace TextGridStyle
)
```
