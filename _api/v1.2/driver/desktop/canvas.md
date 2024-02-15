---
layout: page
tags: [api]
title: Fyne API desktop
---

# desktop
---
```go
import "fyne.io/fyne/driver/desktop"
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
