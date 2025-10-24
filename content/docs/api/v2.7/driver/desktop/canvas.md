---
tags: [api]
title: desktop.Canvas
slug: canvas

aliases:
- /api/v2.7/driver/desktop/canvas

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

#

###

```go
type Canvas interface {
	OnKeyDown() func(*fyne.KeyEvent)
	SetOnKeyDown(func(*fyne.KeyEvent))
	OnKeyUp() func(*fyne.KeyEvent)
	SetOnKeyUp(func(*fyne.KeyEvent))
}
```

Canvas defines the desktop specific extensions to a fyne.Canvas.
