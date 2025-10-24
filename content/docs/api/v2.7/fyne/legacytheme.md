---
tags: [api]
title: fyne.LegacyTheme
slug: legacytheme

aliases:
- /api/v2.7//legacytheme

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type LegacyTheme interface {
	BackgroundColor() color.Color
	ButtonColor() color.Color
	DisabledButtonColor() color.Color
	TextColor() color.Color
	DisabledTextColor() color.Color
	PlaceHolderColor() color.Color
	PrimaryColor() color.Color
	HoverColor() color.Color
	FocusColor() color.Color
	ScrollBarColor() color.Color
	ShadowColor() color.Color

	TextSize() int
	TextFont() Resource
	TextBoldFont() Resource
	TextItalicFont() Resource
	TextBoldItalicFont() Resource
	TextMonospaceFont() Resource

	Padding() int
	IconInlineSize() int
	ScrollBarSize() int
	ScrollBarSmallSize() int
}
```

LegacyTheme defines the requirements of any Fyne theme. This was previously called Theme and is kept for simpler transition of applications built before v2.0.0.


<div class="since">Since: <code>
2.0</code></div>
