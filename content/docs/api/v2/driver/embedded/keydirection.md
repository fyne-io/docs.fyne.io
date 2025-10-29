---
tags: [api]
title: embedded.KeyDirection
slug: keydirection

aliases:
- /api/driver/embedded/keydirection
- /api/driver/embedded/keydirection.html
- /api/v2.0/driver/embedded/keydirection
- /api/v2.0/driver/embedded/keydirection.html
- /api/v2.1/driver/embedded/keydirection
- /api/v2.1/driver/embedded/keydirection.html
- /api/v2.2/driver/embedded/keydirection
- /api/v2.2/driver/embedded/keydirection.html
- /api/v2.3/driver/embedded/keydirection
- /api/v2.3/driver/embedded/keydirection.html
- /api/v2.4/driver/embedded/keydirection
- /api/v2.4/driver/embedded/keydirection.html
- /api/v2.5/driver/embedded/keydirection
- /api/v2.5/driver/embedded/keydirection.html
- /api/v2.6/driver/embedded/keydirection
- /api/v2.6/driver/embedded/keydirection.html
- /api/v2.7/driver/embedded/keydirection
- /api/v2.7/driver/embedded/keydirection.html

package: fyne.io/fyne/v2/driver/embedded
---


---
```go
import "fyne.io/fyne/v2/driver/embedded"
```

## Usage

#### type KeyDirection

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
