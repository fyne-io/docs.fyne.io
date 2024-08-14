---
layout: page
tags: [api]
title: Fyne API "repository.MovableRepository"
package: fyne.io/fyne/v2/storage/repository
---

# repository.MovableRepository
---
```go
import "fyne.io/fyne/v2/storage/repository"
```

## Usage

#### type MovableRepository

```go
type MovableRepository interface {
	Repository

	// Move will be used to implement calls to storage.Move() for the
	// registered scheme of this repository.
	//
	// A generic implementation is provided by GenericMove().
	//
	// NOTE: the first parameter is the source, the second is the
	// destination.
	//
	// NOTE: if storage.Move() is given two URIs of different schemes, it
	// is possible that only the source URI will be of the type this
	// repository is registered to handle. In such cases, implementations
	// are suggested to fail-over to GenericMove().
	//
	// Since: 2.0
	Move(fyne.URI, fyne.URI) error
}
```

MovableRepository is an extension of the Repository interface which also supports moving referenced resources from one URI to another.

Note: both Moveable and Movable are correct spellings, but Movable is newer and more accepted. Source: https://grammarist.com/spelling/movable-moveable/


<div class="since">Since: <code>
2.0</code></div>
