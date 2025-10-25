---
tags: [api]
title: fyne.Settings
slug: settings

aliases:
- /api/v2/fyne/settings.html
- /api/v2.0//settings
- /api/v2.0//settings.html
- /api/v2.1//settings
- /api/v2.1//settings.html
- /api/v2.2//settings
- /api/v2.2//settings.html
- /api/v2.3//settings
- /api/v2.3//settings.html
- /api/v2.4//settings
- /api/v2.4//settings.html
- /api/v2.5//settings
- /api/v2.5//settings.html
- /api/v2.6//settings
- /api/v2.6//settings.html
- /api/v2.7//settings
- /api/v2.7//settings.html

package: fyne.io/fyne/v2
---


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
