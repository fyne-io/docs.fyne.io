---
layout: page
tags: [api]
title: Fyne API "fyne.KeyModifier"
package: fyne.io/fyne/v2
---

# fyne.KeyModifier
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type KeyModifier

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
