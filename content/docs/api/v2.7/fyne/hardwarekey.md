---
tags: [api]
title: fyne.HardwareKey
slug: hardwarekey

aliases:
- /api/v2.7//hardwarekey

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type HardwareKey struct {
	// ScanCode represents a hardware ID for (normally desktop) keyboard events.
	ScanCode int
}
```

HardwareKey contains information associated with physical key events Most applications should use [KeyName] for cross-platform compatibility.
