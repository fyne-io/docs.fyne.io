---
layout: page
tags: [api]
title: Fyne API "repository.CopyableRepository"
---

# repository.CopyableRepository
---
```go
import "fyne.io/fyne/v2/storage/repository"
```

## Usage

#### type CopyableRepository

```go
type CopyableRepository interface {
	Repository

	// Copy will be used to implement calls to storage.Copy() for the
	// registered scheme of this repository.
	//
	// A generic implementation is provided by GenericCopy().
	//
	// NOTE: the first parameter is the source, the second is the
	// destination.
	//
	// NOTE: if storage.Copy() is given two URIs of different schemes, it
	// is possible that only the source URI will be of the type this
	// repository is registered to handle. In such cases, implementations
	// are suggested to fail-over to GenericCopy().
	//
	// Since: 2.0
	Copy(fyne.URI, fyne.URI) error
}
```

CopyableRepository is an extension of the Repository interface which also supports copying referenced resources from one URI to another.


<div class="since">Since: <code>
2.0</code></div>
