---
tags: [api]
title: repository.HierarchicalRepository
slug: hierarchicalrepository

aliases:
- /api/v2/storage/repository/hierarchicalrepository.html
- /api/v2.0/storage/repository/hierarchicalrepository
- /api/v2.0/storage/repository/hierarchicalrepository.html
- /api/v2.1/storage/repository/hierarchicalrepository
- /api/v2.1/storage/repository/hierarchicalrepository.html
- /api/v2.2/storage/repository/hierarchicalrepository
- /api/v2.2/storage/repository/hierarchicalrepository.html
- /api/v2.3/storage/repository/hierarchicalrepository
- /api/v2.3/storage/repository/hierarchicalrepository.html
- /api/v2.4/storage/repository/hierarchicalrepository
- /api/v2.4/storage/repository/hierarchicalrepository.html
- /api/v2.5/storage/repository/hierarchicalrepository
- /api/v2.5/storage/repository/hierarchicalrepository.html
- /api/v2.6/storage/repository/hierarchicalrepository
- /api/v2.6/storage/repository/hierarchicalrepository.html
- /api/v2.7/storage/repository/hierarchicalrepository
- /api/v2.7/storage/repository/hierarchicalrepository.html

package: fyne.io/fyne/v2/storage/repository
---


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
