---
tags: [api]
title: embedded.KeyDirection
slug: keydirection

aliases:
- /api/v2.7/driver/embedded/keydirection

package: fyne.io/fyne/v2/driver/embedded
---


---
```go
import "fyne.io/fyne/v2/driver/embedded"
```

#

###

```go
type KeyDirection uint8
```

KeyDirection specifies the press/release of a key event


<div class="since">Since: <code>
2.7</code></div>

```go
const (
	// KeyPressed specifies that a key was pushed down.
	//
	// Since: 2.7
	KeyPressed KeyDirection = iota

	// KeyReleased indicates a key was let back up.
	//
	// Since: 2.7
	KeyReleased
)
```
