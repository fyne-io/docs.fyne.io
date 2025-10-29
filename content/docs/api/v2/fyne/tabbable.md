---
tags: [api]
title: fyne.Tabbable
slug: tabbable

aliases:
- /api//tabbable
- /api//tabbable.html
- /api/v2.0//tabbable
- /api/v2.0//tabbable.html
- /api/v2.1//tabbable
- /api/v2.1//tabbable.html
- /api/v2.2//tabbable
- /api/v2.2//tabbable.html
- /api/v2.3//tabbable
- /api/v2.3//tabbable.html
- /api/v2.4//tabbable
- /api/v2.4//tabbable.html
- /api/v2.5//tabbable
- /api/v2.5//tabbable.html
- /api/v2.6//tabbable
- /api/v2.6//tabbable.html
- /api/v2.7//tabbable
- /api/v2.7//tabbable.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Tabbable

```go
type Tabbable interface {
	// AcceptsTab is a hook called by the key press handling logic.
	// If it returns true then the Tab key events will be sent using TypedKey.
	AcceptsTab() bool
}
```

Tabbable describes any object that needs to accept the Tab key presses.


<div class="since">Since: <code>
2.1</code></div>
