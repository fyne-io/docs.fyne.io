---
layout: page
tags: [api]
title: Fyne API "fyne.BuildType"
package: fyne.io/fyne/v2
---

# fyne.BuildType
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type BuildType

```go
type BuildType int
```

BuildType defines different modes that an application can be built using.

```go
const (
	// BuildStandard is the normal build mode - it is not debug, test or release mode.
	BuildStandard BuildType = iota
	// BuildDebug is used when a developer would like more information and visual output for app debugging.
	BuildDebug
	// BuildRelease is a final production build, it is like BuildStandard but will use distribution certificates.
	// A release build is typically going to connect to live services and is not usually used during development.
	BuildRelease
)
```
