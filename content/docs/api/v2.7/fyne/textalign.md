---
tags: [api]
title: fyne.TextAlign
slug: textalign

aliases:
- /api/v2.7//textalign

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

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
