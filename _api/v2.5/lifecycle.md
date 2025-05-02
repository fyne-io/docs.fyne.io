---
layout: page
tags: [api]
title: Fyne API "fyne.Lifecycle"
package: fyne.io/fyne/v2
---

# fyne.Lifecycle
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
