---
tags: [api]
title: fyne.Lifecycle
slug: lifecycle

aliases:
- /api/v2/fyne/lifecycle.html
- /api/v2.0//lifecycle
- /api/v2.0//lifecycle.html
- /api/v2.1//lifecycle
- /api/v2.1//lifecycle.html
- /api/v2.2//lifecycle
- /api/v2.2//lifecycle.html
- /api/v2.3//lifecycle
- /api/v2.3//lifecycle.html
- /api/v2.4//lifecycle
- /api/v2.4//lifecycle.html
- /api/v2.5//lifecycle
- /api/v2.5//lifecycle.html
- /api/v2.6//lifecycle
- /api/v2.6//lifecycle.html
- /api/v2.7//lifecycle
- /api/v2.7//lifecycle.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Lifecycle

```go
type Lifecycle interface {
	// SetOnEnteredForeground hooks into the app becoming foreground and gaining focus.
	SetOnEnteredForeground(func())
	// SetOnExitedForeground hooks into the app losing input focus and going into the background.
	SetOnExitedForeground(func())
	// SetOnStarted hooks into an event that says the app is now running.
	SetOnStarted(func())
	// SetOnStopped hooks into an event that says the app is no longer running.
	SetOnStopped(func())
}
```

Lifecycle represents the various phases that an app can transition through.


<div class="since">Since: <code>
2.1</code></div>
