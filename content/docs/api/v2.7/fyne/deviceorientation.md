---
tags: [api]
title: fyne.DeviceOrientation
slug: deviceorientation

aliases:
- /api/v2.7//deviceorientation

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

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
