---
tags: [api]
title: desktop.Canvas
slug: canvas

aliases:
- /api/v2.0/driver/desktop/canvas
- /api/v2.1/driver/desktop/canvas
- /api/v2.2/driver/desktop/canvas
- /api/v2.3/driver/desktop/canvas
- /api/v2.4/driver/desktop/canvas
- /api/v2.5/driver/desktop/canvas
- /api/v2.6/driver/desktop/canvas
- /api/v2.7/driver/desktop/canvas

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type Canvas

```go
type Canvas interface {
	OnKeyDown() func(*fyne.KeyEvent)
	SetOnKeyDown(func(*fyne.KeyEvent))
	OnKeyUp() func(*fyne.KeyEvent)
	SetOnKeyUp(func(*fyne.KeyEvent))
}
```

Canvas defines the desktop specific extensions to a fyne.Canvas.
