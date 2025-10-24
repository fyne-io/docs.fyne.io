---
tags: [api]
title: fyne.CloudProviderStorage
slug: cloudproviderstorage

aliases:
- /api/v2.7//cloudproviderstorage

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type CloudProviderStorage interface {
	// CloudStorage returns a storage provider that will sync documents to the cloud this provider uses.
	CloudStorage(App) Storage
}
```

CloudProviderStorage interface defines the functionality that a cloud provider will include if it is capable of synchronizing user documents.


<div class="since">Since: <code>
2.3</code></div>
