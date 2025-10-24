---
tags: [api]
title: widget.CustomTextGridStyle
slug: customtextgridstyle

aliases:
- /api/v2.7/widget/customtextgridstyle

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type CustomTextGridStyle struct {
	// Since: 2.5
	TextStyle        fyne.TextStyle
	FGColor, BGColor color.Color
}
```

CustomTextGridStyle is a utility type for those not wanting to define their own style types.

###

```go
func (c *CustomTextGridStyle) BackgroundColor() color.Color
```
BackgroundColor is the color a cell should use for the background.

###

```go
func (c *CustomTextGridStyle) Style() fyne.TextStyle
```
Style is the text style a cell should use.

###

```go
func (c *CustomTextGridStyle) TextColor() color.Color
```
TextColor is the color a cell should use for the text.
