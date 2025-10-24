---
tags: [api]
title: repository.CustomURIRepository
slug: customurirepository

aliases:
- /api/v2.0/storage/repository/customurirepository
- /api/v2.1/storage/repository/customurirepository
- /api/v2.2/storage/repository/customurirepository
- /api/v2.3/storage/repository/customurirepository
- /api/v2.4/storage/repository/customurirepository
- /api/v2.5/storage/repository/customurirepository
- /api/v2.6/storage/repository/customurirepository
- /api/v2.7/storage/repository/customurirepository

package: fyne.io/fyne/v2/storage/repository
---


---
```go
import "fyne.io/fyne/v2/storage/repository"
```

## Usage

#### type CustomURIRepository

```go
type CustomURIRepository interface {
	Repository

	// ParseURI will be used to implement calls to storage.ParseURI()
	// for the registered scheme of this repository.
	ParseURI(string) (fyne.URI, error)
}
```

CustomURIRepository is an extension of the repository interface which allows the behavior of storage.ParseURI to be overridden. This is only needed if you wish to generate custom URI types, rather than using Fyne's URI implementation and net/url based parsing.

NOTE: even for URIs with non-RFC3986-compliant encoding, the URI MUST begin with 'scheme:', or storage.ParseURI() will not be able to determine which storage repository to delegate to for parsing.


<div class="since">Since: <code>
2.0</code></div>
