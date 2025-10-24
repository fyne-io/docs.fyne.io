---
tags: [api]
title: fyne.KeyModifier
slug: keymodifier

aliases:
- /api/v2.7//keymodifier

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type KeyModifier int
```

KeyModifier represents any modifier key (shift etc.) that is being pressed together with a key.


<div class="since">Since: <code>
2.2</code></div>

```go
const (
	// KeyModifierShift represents a shift key being held
	//
	// Since: 2.2
	KeyModifierShift KeyModifier = 1 << iota
	// KeyModifierControl represents the ctrl key being held
	//
	// Since: 2.2
	KeyModifierControl
	// KeyModifierAlt represents either alt keys being held
	//
	// Since: 2.2
	KeyModifierAlt
	// KeyModifierSuper represents either super keys being held
	//
	// Since: 2.2
	KeyModifierSuper
)
```
