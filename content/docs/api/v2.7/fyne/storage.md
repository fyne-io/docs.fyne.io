---
tags: [api]
title: fyne.Storage
slug: storage

aliases:
- /api/v2.7//storage

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Storage interface {
	RootURI() URI

	Create(name string) (URIWriteCloser, error)
	Open(name string) (URIReadCloser, error)
	Save(name string) (URIWriteCloser, error)
	Remove(name string) error

	List() []string
}
```

Storage is used to manage file storage inside an application sandbox. The files managed by this interface are unique to the current application.
