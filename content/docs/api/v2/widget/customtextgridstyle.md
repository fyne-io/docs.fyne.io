---
tags: [api]
title: widget.CustomTextGridStyle
slug: customtextgridstyle

aliases:
- /api/v2/widget/customtextgridstyle.html
- /api/v2.0/widget/customtextgridstyle
- /api/v2.0/widget/customtextgridstyle.html
- /api/v2.1/widget/customtextgridstyle
- /api/v2.1/widget/customtextgridstyle.html
- /api/v2.2/widget/customtextgridstyle
- /api/v2.2/widget/customtextgridstyle.html
- /api/v2.3/widget/customtextgridstyle
- /api/v2.3/widget/customtextgridstyle.html
- /api/v2.4/widget/customtextgridstyle
- /api/v2.4/widget/customtextgridstyle.html
- /api/v2.5/widget/customtextgridstyle
- /api/v2.5/widget/customtextgridstyle.html
- /api/v2.6/widget/customtextgridstyle
- /api/v2.6/widget/customtextgridstyle.html
- /api/v2.7/widget/customtextgridstyle
- /api/v2.7/widget/customtextgridstyle.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type CustomTextGridStyle

```go
type CustomTextGridStyle struct {
	// Since: 2.5
	TextStyle        fyne.TextStyle
	FGColor, BGColor color.Color
}
```

CustomTextGridStyle is a utility type for those not wanting to define their own style types.

#### func (*CustomTextGridStyle) BackgroundColor

```go
func (c *CustomTextGridStyle) BackgroundColor() color.Color
```
BackgroundColor is the color a cell should use for the background.

#### func (*CustomTextGridStyle) Style

```go
func (c *CustomTextGridStyle) Style() fyne.TextStyle
```
Style is the text style a cell should use.

#### func (*CustomTextGridStyle) TextColor

```go
func (c *CustomTextGridStyle) TextColor() color.Color
```
TextColor is the color a cell should use for the text.
