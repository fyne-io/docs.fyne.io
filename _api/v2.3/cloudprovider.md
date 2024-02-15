---
layout: page
tags: [api]
title: Fyne API "fyne.CloudProvider"
---

# fyne.CloudProvider
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type CloudProvider

```go
type CloudProvider interface {
	// ProviderDescription returns a more detailed description of this cloud provider.
	ProviderDescription() string
	// ProviderIcon returns an icon resource that is associated with the given cloud service.
	ProviderIcon() Resource
	// ProviderName returns the name of this cloud provider, usually the name of the service it uses.
	ProviderName() string

	// Cleanup is called when this provider is no longer used and should be disposed.
	// This is guaranteed to execute before a new provider is `Setup`
	Cleanup(App)
	// Setup is called when this provider is being used for the first time.
	// Returning an error will exit the cloud setup process, though it can be retried.
	Setup(App) error
}
```

CloudProvider specifies the identifying information of a cloud provider. This information is mostly used by the `fyne.io/cloud ShowSettings' user flow.


<div class="since">Since: <code>
2.3</code></div>
