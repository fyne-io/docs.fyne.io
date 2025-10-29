---
tags: [api]
title: mobile.Touchable
slug: touchable

aliases:
- /api/driver/mobile/touchable
- /api/driver/mobile/touchable.html
- /api/v2.0/driver/mobile/touchable
- /api/v2.0/driver/mobile/touchable.html
- /api/v2.1/driver/mobile/touchable
- /api/v2.1/driver/mobile/touchable.html
- /api/v2.2/driver/mobile/touchable
- /api/v2.2/driver/mobile/touchable.html
- /api/v2.3/driver/mobile/touchable
- /api/v2.3/driver/mobile/touchable.html
- /api/v2.4/driver/mobile/touchable
- /api/v2.4/driver/mobile/touchable.html
- /api/v2.5/driver/mobile/touchable
- /api/v2.5/driver/mobile/touchable.html
- /api/v2.6/driver/mobile/touchable
- /api/v2.6/driver/mobile/touchable.html
- /api/v2.7/driver/mobile/touchable
- /api/v2.7/driver/mobile/touchable.html

package: fyne.io/fyne/v2/driver/mobile
---


---
```go
import "fyne.io/fyne/v2/driver/mobile"
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
