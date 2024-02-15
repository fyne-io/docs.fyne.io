---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type TextGridCell

```go
type TextGridCell struct {
	Rune  rune
	Style TextGridStyle
}
```

TextGridCell represents a single cell in a text grid. It has a rune for the text content and a style associated with it.
