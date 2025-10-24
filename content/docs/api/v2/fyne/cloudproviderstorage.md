---
tags: [api]
title: fyne.CloudProviderStorage
slug: cloudproviderstorage

aliases:
- /api/v2.0//cloudproviderstorage
- /api/v2.1//cloudproviderstorage
- /api/v2.2//cloudproviderstorage
- /api/v2.3//cloudproviderstorage
- /api/v2.4//cloudproviderstorage
- /api/v2.5//cloudproviderstorage
- /api/v2.6//cloudproviderstorage
- /api/v2.7//cloudproviderstorage

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type CloudProviderStorage

```go
type CloudProviderStorage interface {
	// CloudStorage returns a storage provider that will sync documents to the cloud this provider uses.
	CloudStorage(App) Storage
}
```

CloudProviderStorage interface defines the functionality that a cloud provider will include if it is capable of synchronizing user documents.


<div class="since">Since: <code>
2.3</code></div>
