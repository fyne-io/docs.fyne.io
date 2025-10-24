---
tags: [api]
title: widget.TextGridRow
slug: textgridrow

aliases:
- /api/v2.0/widget/textgridrow
- /api/v2.1/widget/textgridrow
- /api/v2.2/widget/textgridrow
- /api/v2.3/widget/textgridrow
- /api/v2.4/widget/textgridrow
- /api/v2.5/widget/textgridrow
- /api/v2.6/widget/textgridrow
- /api/v2.7/widget/textgridrow

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type TextGridRow

```go
type TextGridRow struct {
	Cells []TextGridCell
	Style TextGridStyle
}
```

TextGridRow represents a row of cells cell in a text grid. It contains the cells for the row and an optional style.
