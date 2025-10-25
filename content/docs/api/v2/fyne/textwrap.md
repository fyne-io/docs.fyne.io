---
tags: [api]
title: fyne.TextWrap
slug: textwrap

aliases:
- /api/v2/fyne/textwrap.html
- /api/v2.0//textwrap
- /api/v2.0//textwrap.html
- /api/v2.1//textwrap
- /api/v2.1//textwrap.html
- /api/v2.2//textwrap
- /api/v2.2//textwrap.html
- /api/v2.3//textwrap
- /api/v2.3//textwrap.html
- /api/v2.4//textwrap
- /api/v2.4//textwrap.html
- /api/v2.5//textwrap
- /api/v2.5//textwrap.html
- /api/v2.6//textwrap
- /api/v2.6//textwrap.html
- /api/v2.7//textwrap
- /api/v2.7//textwrap.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type TextWrap

```go
type TextWrap int
```

TextWrap represents how text longer than the widget's width will be wrapped.

```go
const (
	// TextWrapOff extends the widget's width to fit the text, no wrapping is applied.
	TextWrapOff TextWrap = iota
	// TextTruncate trims the text to the widget's width, no wrapping is applied.
	// If an entry is asked to truncate it will provide scrolling capabilities.
	//
	// Deprecated: Use [TextTruncateClip] value of the widget `Truncation` field instead
	TextTruncate
	// TextWrapBreak trims the line of characters to the widget's width adding the excess as new line.
	// An Entry with text wrapping will scroll vertically if there is not enough space for all the text.
	TextWrapBreak
	// TextWrapWord trims the line of words to the widget's width adding the excess as new line.
	// An Entry with text wrapping will scroll vertically if there is not enough space for all the text.
	TextWrapWord
)
```
