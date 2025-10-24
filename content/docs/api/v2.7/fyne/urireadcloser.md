---
tags: [api]
title: fyne.URIReadCloser
slug: urireadcloser

aliases:
- /api/v2.7//urireadcloser

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type URIReadCloser interface {
	io.ReadCloser

	URI() URI
}
```

URIReadCloser represents a cross platform data stream from a file or provider of data. It may refer to an item on a filesystem or data in another application that we have access to.
