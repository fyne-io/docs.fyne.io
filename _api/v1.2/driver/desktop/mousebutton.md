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

#### type MouseButton

```go
type MouseButton int
```

MouseButton represents a single button in a desktop MouseEvent

```go
const (
	// LeftMouseButton is the most common mouse button - on some systems the only one
	LeftMouseButton MouseButton = iota + 1

	// RightMouseButton is the secondary button on most mouse input devices.
	RightMouseButton
)
```
