---
tags: [api]
title: widget.TextGridCell
slug: textgridcell

aliases:
- /api/v2.0/widget/textgridcell
- /api/v2.1/widget/textgridcell
- /api/v2.2/widget/textgridcell
- /api/v2.3/widget/textgridcell
- /api/v2.4/widget/textgridcell
- /api/v2.5/widget/textgridcell
- /api/v2.6/widget/textgridcell
- /api/v2.7/widget/textgridcell

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
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
