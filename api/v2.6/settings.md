---
layout: page
tags: [api]
title: Fyne API "fyne.Settings"
package: fyne.io/fyne/v2
---

# fyne.Settings
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Settings

```go
type Settings interface {
	Theme() Theme
	SetTheme(Theme)
	// ThemeVariant defines which preferred version of a theme should be used (i.e. light or dark)
	//
	// Since: 2.0
	ThemeVariant() ThemeVariant
	Scale() float32
	// PrimaryColor indicates a user preference for a named primary color
	//
	// Since: 1.4
	PrimaryColor() string

	// AddChangeListener subscribes to settings change events over a channel.
	//
	// Deprecated: Use AddListener instead, which uses a callback-based API
	// with the callback guaranteed to be invoked on the app goroutine.
	AddChangeListener(chan Settings)

	// AddListener registers a callback that is invoked whenever the settings change.
	//
	// Since: 2.6
	AddListener(func(Settings))

	BuildType() BuildType

	ShowAnimations() bool
}
```

Settings describes the application configuration available.
