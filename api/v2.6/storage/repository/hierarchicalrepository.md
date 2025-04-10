---
layout: page
tags: [api]
title: Fyne API "repository.HierarchicalRepository"
package: fyne.io/fyne/v2/storage/repository
---

# repository.HierarchicalRepository
---
```go
import "fyne.io/fyne/v2/storage/repository"
```

## Usage

#### type HierarchicalRepository

```go
type HierarchicalRepository interface {
	Repository

	// Parent will be used to implement calls to storage.Parent() for the
	// registered scheme of this repository.
	//
	// A generic implementation is provided in GenericParent(), which
	// is based on the RFC3986 definition of a URI parent.
	//
	// Since: 2.0
	Parent(fyne.URI) (fyne.URI, error)

	// Child will be used to implement calls to storage.Child() for
	// the registered scheme of this repository.
	//
	// A generic implementation is provided in GenericParent(), which
	// is based on RFC3986.
	//
	// Since: 2.0
	Child(fyne.URI, string) (fyne.URI, error)
}
```

HierarchicalRepository is an extension of the Repository interface which also supports determining the parent and child items of a URI.


<div class="since">Since: <code>
2.0</code></div>
