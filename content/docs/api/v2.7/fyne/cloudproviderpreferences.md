---
tags: [api]
title: fyne.CloudProviderPreferences
slug: cloudproviderpreferences

aliases:
- /api/v2.7//cloudproviderpreferences

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type CloudProviderPreferences interface {
	// CloudPreferences returns a preference provider that will sync values to the cloud this provider uses.
	CloudPreferences(App) Preferences
}
```

CloudProviderPreferences interface defines the functionality that a cloud provider will include if it is capable of synchronizing user preferences.


<div class="since">Since: <code>
2.3</code></div>
