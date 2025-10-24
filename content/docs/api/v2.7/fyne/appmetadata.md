---
tags: [api]
title: fyne.AppMetadata
slug: appmetadata

aliases:
- /api/v2.7//appmetadata

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

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
	// Release if true this binary was build in release mode
	// Since: 2.3
	Release bool
	// Custom contain the custom metadata defined either in FyneApp.toml or on the compile command line
	// Since: 2.3
	Custom map[string]string
	// Migrations allows an app to opt into features before they are standard
	// Since: 2.6
	Migrations map[string]bool
}
```

AppMetadata captures the build metadata for an application.


<div class="since">Since: <code>
2.2</code></div>
