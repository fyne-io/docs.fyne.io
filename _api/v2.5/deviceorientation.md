---
layout: page
tags: [api]
title: Fyne API "fyne.DeviceOrientation"
package: fyne.io/fyne/v2
---

# fyne.DeviceOrientation
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type DeviceOrientation

```go
type DeviceOrientation int
```

DeviceOrientation represents the different ways that a mobile device can be held

```go
const (
	// OrientationVertical is the default vertical orientation
	OrientationVertical DeviceOrientation = iota
	// OrientationVerticalUpsideDown is the portrait orientation held upside down
	OrientationVerticalUpsideDown
	// OrientationHorizontalLeft is used to indicate a landscape orientation with the top to the left
	OrientationHorizontalLeft
	// OrientationHorizontalRight is used to indicate a landscape orientation with the top to the right
	OrientationHorizontalRight
)
```
