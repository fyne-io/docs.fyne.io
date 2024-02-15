---
layout: page
tags: [api]
title: Fyne API "fyne.TextWrap"
---

# fyne.TextWrap
---
```go
import "fyne.io/fyne"
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
	TextTruncate
	// TextWrapBreak trims the line of characters to the widget's width adding the excess as new line.
	TextWrapBreak
	// TextWrapWord trims the line of words to the widget's width adding the excess as new line.
	TextWrapWord
)
```
