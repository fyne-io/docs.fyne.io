---
layout: page
tags: [api]
title: Fyne API "repository.CustomURIRepository"
---

# repository.CustomURIRepository
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
