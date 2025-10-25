---
tags: [api]
title: fyne.DeviceOrientation
slug: deviceorientation

aliases:
- /api/v2/fyne/deviceorientation.html
- /api/v2.0//deviceorientation
- /api/v2.0//deviceorientation.html
- /api/v2.1//deviceorientation
- /api/v2.1//deviceorientation.html
- /api/v2.2//deviceorientation
- /api/v2.2//deviceorientation.html
- /api/v2.3//deviceorientation
- /api/v2.3//deviceorientation.html
- /api/v2.4//deviceorientation
- /api/v2.4//deviceorientation.html
- /api/v2.5//deviceorientation
- /api/v2.5//deviceorientation.html
- /api/v2.6//deviceorientation
- /api/v2.6//deviceorientation.html
- /api/v2.7//deviceorientation
- /api/v2.7//deviceorientation.html

package: fyne.io/fyne/v2
---


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
