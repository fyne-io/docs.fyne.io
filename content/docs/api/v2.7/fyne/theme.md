---
tags: [api]
title: fyne.Theme
slug: theme

aliases:
- /api/v2.7//theme

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Theme interface {
	Color(ThemeColorName, ThemeVariant) color.Color
	Font(TextStyle) Resource
	Icon(ThemeIconName) Resource
	Size(ThemeSizeName) float32
}
```

Theme defines the method to look up colors, sizes and fonts that make up a Fyne theme.


<div class="since">Since: <code>
2.0</code></div>
