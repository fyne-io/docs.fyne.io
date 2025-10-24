---
tags: [api]
title: widget.TextGridStyle
slug: textgridstyle

aliases:
- /api/v2.7/widget/textgridstyle

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

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
