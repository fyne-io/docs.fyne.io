---
tags: [api]
title: fyne.CloudProviderPreferences
slug: cloudproviderpreferences

aliases:
- /api/v2.0//cloudproviderpreferences
- /api/v2.1//cloudproviderpreferences
- /api/v2.2//cloudproviderpreferences
- /api/v2.3//cloudproviderpreferences
- /api/v2.4//cloudproviderpreferences
- /api/v2.5//cloudproviderpreferences
- /api/v2.6//cloudproviderpreferences
- /api/v2.7//cloudproviderpreferences

package: fyne.io/fyne/v2
---


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
