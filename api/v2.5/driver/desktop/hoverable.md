---
layout: page
tags: [api]
title: Fyne API "desktop.Hoverable"
package: fyne.io/fyne/v2/driver/desktop
---

# desktop.Hoverable
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type Hoverable

```go
type Hoverable interface {
	// MouseIn is a hook that is called if the mouse pointer enters the element.
	MouseIn(*MouseEvent)
	// MouseMoved is a hook that is called if the mouse pointer moved over the element.
	MouseMoved(*MouseEvent)
	// MouseOut is a hook that is called if the mouse pointer leaves the element.
	MouseOut()
}
```

Hoverable is used when a canvas object wishes to know if a pointer device moves over it.
