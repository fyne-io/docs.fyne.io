---
layout: page
tags: [api]
title: Fyne API "desktop.MouseButton"
package: fyne.io/fyne/v2/driver/desktop
---

# desktop.MouseButton
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
