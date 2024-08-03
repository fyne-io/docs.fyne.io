---
layout: page
tags: [api]
title: Fyne API "fyne.Disableable"
package: fyne.io/fyne/v2
---

# fyne.Disableable
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

Disableable describes any CanvasObject that can be disabled. This is primarily used with objects that also implement the Tappable interface.
