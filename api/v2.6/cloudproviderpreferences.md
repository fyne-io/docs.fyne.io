---
layout: page
tags: [api]
title: Fyne API "fyne.CloudProviderPreferences"
package: fyne.io/fyne/v2
---

# fyne.CloudProviderPreferences
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type CloudProviderPreferences

```go
type CloudProviderPreferences interface {
	// CloudPreferences returns a preference provider that will sync values to the cloud this provider uses.
	CloudPreferences(App) Preferences
}
```

CloudProviderPreferences interface defines the functionality that a cloud provider will include if it is capable of synchronizing user preferences.


<div class="since">Since: <code>
2.3</code></div>
