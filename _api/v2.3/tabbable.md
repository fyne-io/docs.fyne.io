---
layout: page
tags: [api]
title: Fyne API "fyne.Tabbable"
---

# fyne.Tabbable
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Tabbable

```go
type Tabbable interface {
	// AcceptsTab() is a hook called by the key press handling logic.
	// If it returns true then the Tab key events will be sent using TypedKey.
	AcceptsTab() bool
}
```

Tabbable describes any object that needs to accept the Tab key presses.


<div class="since">Since: <code>
2.1</code></div>
