---
tags: [api]
title: fyne.TextAlign
slug: textalign

aliases:
- /api/v2/fyne/textalign.html
- /api/v2.0//textalign
- /api/v2.0//textalign.html
- /api/v2.1//textalign
- /api/v2.1//textalign.html
- /api/v2.2//textalign
- /api/v2.2//textalign.html
- /api/v2.3//textalign
- /api/v2.3//textalign.html
- /api/v2.4//textalign
- /api/v2.4//textalign.html
- /api/v2.5//textalign
- /api/v2.5//textalign.html
- /api/v2.6//textalign
- /api/v2.6//textalign.html
- /api/v2.7//textalign
- /api/v2.7//textalign.html

package: fyne.io/fyne/v2
---


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
