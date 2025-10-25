---
tags: [api]
title: repository.ListableRepository
slug: listablerepository

aliases:
- /api/v2/storage/repository/listablerepository.html
- /api/v2.0/storage/repository/listablerepository
- /api/v2.0/storage/repository/listablerepository.html
- /api/v2.1/storage/repository/listablerepository
- /api/v2.1/storage/repository/listablerepository.html
- /api/v2.2/storage/repository/listablerepository
- /api/v2.2/storage/repository/listablerepository.html
- /api/v2.3/storage/repository/listablerepository
- /api/v2.3/storage/repository/listablerepository.html
- /api/v2.4/storage/repository/listablerepository
- /api/v2.4/storage/repository/listablerepository.html
- /api/v2.5/storage/repository/listablerepository
- /api/v2.5/storage/repository/listablerepository.html
- /api/v2.6/storage/repository/listablerepository
- /api/v2.6/storage/repository/listablerepository.html
- /api/v2.7/storage/repository/listablerepository
- /api/v2.7/storage/repository/listablerepository.html

package: fyne.io/fyne/v2/storage/repository
---


---
```go
import "fyne.io/fyne/v2/storage/repository"
```

## Usage

#### type ListableRepository

```go
type ListableRepository interface {
	Repository

	// CanList will be used to implement calls to storage.Listable() for
	// the registered scheme of this repository.
	//
	// Since: 2.0
	CanList(u fyne.URI) (bool, error)

	// List will be used to implement calls to storage.List() for the
	// registered scheme of this repository.
	//
	// Since: 2.0
	List(u fyne.URI) ([]fyne.URI, error)

	// CreateListable will be used to implement calls to
	// storage.CreateListable() for the registered scheme of this
	// repository.
	//
	// Since: 2.0
	CreateListable(u fyne.URI) error
}
```

ListableRepository is an extension of the Repository interface which also supports obtaining directory listings (generally analogous to a directory listing) for URIs of the scheme it is registered to.


<div class="since">Since: <code>
2.0</code></div>
