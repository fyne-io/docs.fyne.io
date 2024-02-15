---
layout: page
tags: [api]
title: Fyne API "fyne.Settings"
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

	AddChangeListener(chan Settings)
	BuildType() BuildType
}
```

Settings describes the application configuration available.
