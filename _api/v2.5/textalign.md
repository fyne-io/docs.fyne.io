---
layout: page
tags: [api]
title: Fyne API "fyne.TextAlign"
package: fyne.io/fyne/v2
---

# fyne.TextAlign
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type TextAlign

```go
type TextAlign int
```

TextAlign represents the horizontal alignment of text within a widget or canvas object.

```go
const (
	// TextAlignLeading specifies a left alignment for left-to-right languages.
	TextAlignLeading TextAlign = iota
	// TextAlignCenter places the text centrally within the available space.
	TextAlignCenter
	// TextAlignTrailing will align the text right for a left-to-right language.
	TextAlignTrailing
)
```
