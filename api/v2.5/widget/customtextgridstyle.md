---
layout: page
tags: [api]
title: Fyne API "widget.CustomTextGridStyle"
package: fyne.io/fyne/v2/widget
---

# widget.CustomTextGridStyle
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
