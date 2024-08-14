---
layout: page
tags: [api]
title: Fyne API "repository.WritableRepository"
package: fyne.io/fyne/v2/storage/repository
---

# repository.WritableRepository
---
```go
import "fyne.io/fyne/v2/storage/repository"
```

## Usage

#### type WritableRepository

```go
type WritableRepository interface {
	Repository

	// Writer will be used to implement calls to storage.WriterTo() for
	// the registered scheme of this repository.
	//
	// Since: 2.0
	Writer(u fyne.URI) (fyne.URIWriteCloser, error)

	// CanWrite will be used to implement calls to storage.CanWrite() for
	// the registered scheme of this repository.
	//
	// Since: 2.0
	CanWrite(u fyne.URI) (bool, error)

	// Delete will be used to implement calls to storage.Delete() for the
	// registered scheme of this repository.
	//
	// Since: 2.0
	Delete(u fyne.URI) error
}
```

WritableRepository is an extension of the Repository interface which also supports obtaining a writer for URIs of the scheme it is registered to.


<div class="since">Since: <code>
2.0</code></div>
