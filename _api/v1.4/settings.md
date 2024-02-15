---
layout: page
tags: [api]
title: Fyne API "fyne.Settings"
---

# fyne.Settings
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Settings

```go
type Settings interface {
	Theme() Theme
	SetTheme(Theme)
	Scale() float32
	PrimaryColor() string

	AddChangeListener(chan Settings)
	BuildType() BuildType
}
```

Settings describes the application configuration available.
