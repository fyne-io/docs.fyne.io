---
layout: page
tags: [api]
title: Fyne API "mobile.Touchable"
---

# mobile.Touchable
---
```go
import "fyne.io/fyne/driver/mobile"
```

## Usage

#### type Touchable

```go
type Touchable interface {
	TouchDown(*TouchEvent)
	TouchUp(*TouchEvent)
	TouchCancel(*TouchEvent)
}
```

Touchable represents mobile touch events that can be sent to CanvasObjects
