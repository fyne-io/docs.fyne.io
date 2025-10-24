---
tags: [api]
title: desktop.MouseButton
slug: mousebutton

aliases:
- /api/v2.0/driver/desktop/mousebutton
- /api/v2.1/driver/desktop/mousebutton
- /api/v2.2/driver/desktop/mousebutton
- /api/v2.3/driver/desktop/mousebutton
- /api/v2.4/driver/desktop/mousebutton
- /api/v2.5/driver/desktop/mousebutton
- /api/v2.6/driver/desktop/mousebutton
- /api/v2.7/driver/desktop/mousebutton

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type MouseButton

```go
type MouseButton int
```

MouseButton represents a single button in a desktop MouseEvent

```go
const (
	// MouseButtonPrimary is the most common mouse button - on some systems the only one.
	// This will normally be on the left side of a mouse.
	//
	// Since: 2.0
	MouseButtonPrimary MouseButton = 1 << iota

	// MouseButtonSecondary is the secondary button on most mouse input devices.
	// This will normally be on the right side of a mouse.
	//
	// Since: 2.0
	MouseButtonSecondary

	// MouseButtonTertiary is the middle button on the mouse, assuming it has one.
	//
	// Since: 2.0
	MouseButtonTertiary

	// LeftMouseButton is the most common mouse button - on some systems the only one.
	//
	// Deprecated: use MouseButtonPrimary which will adapt to mouse configuration.
	LeftMouseButton = MouseButtonPrimary

	// RightMouseButton is the secondary button on most mouse input devices.
	//
	// Deprecated: use MouseButtonSecondary which will adapt to mouse configuration.
	RightMouseButton = MouseButtonSecondary
)
```
