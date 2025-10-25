---
tags: [api]
title: fyne.Clipboard
slug: clipboard

aliases:
- /api/v2/fyne/clipboard.html
- /api/v2.0//clipboard
- /api/v2.0//clipboard.html
- /api/v2.1//clipboard
- /api/v2.1//clipboard.html
- /api/v2.2//clipboard
- /api/v2.2//clipboard.html
- /api/v2.3//clipboard
- /api/v2.3//clipboard.html
- /api/v2.4//clipboard
- /api/v2.4//clipboard.html
- /api/v2.5//clipboard
- /api/v2.5//clipboard.html
- /api/v2.6//clipboard
- /api/v2.6//clipboard.html
- /api/v2.7//clipboard
- /api/v2.7//clipboard.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Clipboard

```go
type Clipboard interface {
	// Content returns the clipboard content
	Content() string
	// SetContent sets the clipboard content
	SetContent(content string)
}
```

Clipboard represents the system clipboard interface
