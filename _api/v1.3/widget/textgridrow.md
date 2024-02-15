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

#### type TextGridRow

```go
type TextGridRow struct {
	Cells []TextGridCell
	Style TextGridStyle
}
```

TextGridRow represents a row of cells cell in a text grid. It contains the cells for the row and an optional style.
