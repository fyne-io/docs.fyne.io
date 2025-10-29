---
tags: [api]
title: fyne.BuildType
slug: buildtype

aliases:
- /api//buildtype
- /api//buildtype.html
- /api/v2.0//buildtype
- /api/v2.0//buildtype.html
- /api/v2.1//buildtype
- /api/v2.1//buildtype.html
- /api/v2.2//buildtype
- /api/v2.2//buildtype.html
- /api/v2.3//buildtype
- /api/v2.3//buildtype.html
- /api/v2.4//buildtype
- /api/v2.4//buildtype.html
- /api/v2.5//buildtype
- /api/v2.5//buildtype.html
- /api/v2.6//buildtype
- /api/v2.6//buildtype.html
- /api/v2.7//buildtype
- /api/v2.7//buildtype.html

package: fyne.io/fyne/v2
---


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
	// BuildRelease is a final production build, it is like [BuildStandard] but will use distribution certificates.
	// A release build is typically going to connect to live services and is not usually used during development.
	BuildRelease
)
```
