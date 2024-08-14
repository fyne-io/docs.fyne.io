---
layout: page
tags: [api]
title: Fyne API "fyne.TextWrap"
package: fyne.io/fyne/v2
---

# fyne.TextWrap
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
	// Deprecated: Use `TextTruncateClip` value of the widget `Truncation` field instead
	TextTruncate
	// TextWrapBreak trims the line of characters to the widget's width adding the excess as new line.
	// An Entry with text wrapping will scroll vertically if there is not enough space for all the text.
	TextWrapBreak
	// TextWrapWord trims the line of words to the widget's width adding the excess as new line.
	// An Entry with text wrapping will scroll vertically if there is not enough space for all the text.
	TextWrapWord
)
```
