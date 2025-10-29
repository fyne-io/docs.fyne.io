---
tags: [api]
title: widget.RichTextStyle
slug: richtextstyle

aliases:
- /api/widget/richtextstyle
- /api/widget/richtextstyle.html
- /api/v2.0/widget/richtextstyle
- /api/v2.0/widget/richtextstyle.html
- /api/v2.1/widget/richtextstyle
- /api/v2.1/widget/richtextstyle.html
- /api/v2.2/widget/richtextstyle
- /api/v2.2/widget/richtextstyle.html
- /api/v2.3/widget/richtextstyle
- /api/v2.3/widget/richtextstyle.html
- /api/v2.4/widget/richtextstyle
- /api/v2.4/widget/richtextstyle.html
- /api/v2.5/widget/richtextstyle
- /api/v2.5/widget/richtextstyle.html
- /api/v2.6/widget/richtextstyle
- /api/v2.6/widget/richtextstyle.html
- /api/v2.7/widget/richtextstyle
- /api/v2.7/widget/richtextstyle.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type RichTextStyle

```go
type RichTextStyle struct {
	Alignment fyne.TextAlign
	ColorName fyne.ThemeColorName
	Inline    bool
	SizeName  fyne.ThemeSizeName // The theme name of the text size to use, if blank will be the standard text size
	TextStyle fyne.TextStyle
}
```

RichTextStyle describes the details of a text object inside a RichText widget.


<div class="since">Since: <code>
2.1</code></div>
