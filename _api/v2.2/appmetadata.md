---
layout: page
tags: [api]
title: Fyne API "fyne.AppMetadata"
---

# fyne.AppMetadata
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type AppMetadata

```go
type AppMetadata struct {
	// ID is the unique ID of this application, used by many distribution platforms.
	ID string
	// Name is the human friendly name of this app.
	Name string
	// Version represents the version of this application, normally following semantic versioning.
	Version string
	// Build is the build number of this app, some times appended to the version number.
	Build int
	// Icon contains, if present, a resource of the icon that was bundled at build time.
	Icon Resource
}
```

AppMetadata captures the build metadata for an application.


<div class="since">Since: <code>
2.2</code></div>
