---
tags: [api]
title: fyne.URIWriteCloser
slug: uriwritecloser

aliases:
- /api/v2.7//uriwritecloser

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type URIWriteCloser interface {
	io.WriteCloser

	URI() URI
}
```

URIWriteCloser represents a cross platform data writer for a file resource. This will normally refer to a local file resource.
