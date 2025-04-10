---
layout: page
tags: [api]
title: Fyne API "fyne.CloudProviderStorage"
package: fyne.io/fyne/v2
---

# fyne.CloudProviderStorage
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
