---
tags: [api]
title: fyne.Theme
slug: theme

aliases:
- /api//theme
- /api//theme.html
- /api/v2.0//theme
- /api/v2.0//theme.html
- /api/v2.1//theme
- /api/v2.1//theme.html
- /api/v2.2//theme
- /api/v2.2//theme.html
- /api/v2.3//theme
- /api/v2.3//theme.html
- /api/v2.4//theme
- /api/v2.4//theme.html
- /api/v2.5//theme
- /api/v2.5//theme.html
- /api/v2.6//theme
- /api/v2.6//theme.html
- /api/v2.7//theme
- /api/v2.7//theme.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Theme

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
