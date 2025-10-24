---
tags: [api]
title: repository.AppendableRepository
slug: appendablerepository

aliases:
- /api/v2.7/storage/repository/appendablerepository

package: fyne.io/fyne/v2/storage/repository
---


---
```go
import "fyne.io/fyne/v2/storage/repository"
```

#

###

```go
type AppendableRepository interface {
	WritableRepository

	// Appender will be used to call a Writer without truncating the
	// file if it exists
	//
	// Since: 2.6
	Appender(u fyne.URI) (fyne.URIWriteCloser, error)
}
```

AppendableRepository is an extension of the WritableRepository interface which also supports opening a writer for URIs in append mode, without truncating their contents


<div class="since">Since: <code>
2.6</code></div>
