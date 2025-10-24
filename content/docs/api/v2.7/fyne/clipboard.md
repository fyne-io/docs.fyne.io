---
tags: [api]
title: fyne.Clipboard
slug: clipboard

aliases:
- /api/v2.7//clipboard

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Clipboard interface {
	// Content returns the clipboard content
	Content() string
	// SetContent sets the clipboard content
	SetContent(content string)
}
```

Clipboard represents the system clipboard interface
