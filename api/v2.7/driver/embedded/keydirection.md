---
layout: page
tags: [api]
title: Fyne API "embedded.KeyDirection"
package: fyne.io/fyne/v2/driver/embedded
---

# embedded.KeyDirection
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
