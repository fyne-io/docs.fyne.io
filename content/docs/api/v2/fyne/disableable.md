---
tags: [api]
title: fyne.Disableable
slug: disableable

aliases:
- /api/v2/fyne/disableable.html
- /api/v2.0//disableable
- /api/v2.0//disableable.html
- /api/v2.1//disableable
- /api/v2.1//disableable.html
- /api/v2.2//disableable
- /api/v2.2//disableable.html
- /api/v2.3//disableable
- /api/v2.3//disableable.html
- /api/v2.4//disableable
- /api/v2.4//disableable.html
- /api/v2.5//disableable
- /api/v2.5//disableable.html
- /api/v2.6//disableable
- /api/v2.6//disableable.html
- /api/v2.7//disableable
- /api/v2.7//disableable.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Disableable

```go
type Disableable interface {
	Enable()
	Disable()
	Disabled() bool
}
```

Disableable describes any [CanvasObject] that can be disabled. This is primarily used with objects that also implement the Tappable interface.
