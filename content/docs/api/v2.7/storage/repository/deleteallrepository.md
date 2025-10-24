---
tags: [api]
title: repository.DeleteAllRepository
slug: deleteallrepository

aliases:
- /api/v2.7/storage/repository/deleteallrepository

package: fyne.io/fyne/v2/storage/repository
---


---
```go
import "fyne.io/fyne/v2/storage/repository"
```

#

###

```go
type DeleteAllRepository interface {
	WritableRepository

	// DeleteAll will be used to implement calls to storage.DeleteAll() for the
	// registered scheme of this repository.
	//
	// A generic implementation is provided by GenericDeleteAll().
	//
	// Since: 2.7
	DeleteAll(fyne.URI) error
}
```

DeleteAllRepository is an extension of the WritableRepository interface which also supports deleting a URI and all its children.


<div class="since">Since: <code>
2.7</code></div>
