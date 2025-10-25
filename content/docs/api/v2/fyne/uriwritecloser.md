---
tags: [api]
title: fyne.URIWriteCloser
slug: uriwritecloser

aliases:
- /api/v2/fyne/uriwritecloser.html
- /api/v2.0//uriwritecloser
- /api/v2.0//uriwritecloser.html
- /api/v2.1//uriwritecloser
- /api/v2.1//uriwritecloser.html
- /api/v2.2//uriwritecloser
- /api/v2.2//uriwritecloser.html
- /api/v2.3//uriwritecloser
- /api/v2.3//uriwritecloser.html
- /api/v2.4//uriwritecloser
- /api/v2.4//uriwritecloser.html
- /api/v2.5//uriwritecloser
- /api/v2.5//uriwritecloser.html
- /api/v2.6//uriwritecloser
- /api/v2.6//uriwritecloser.html
- /api/v2.7//uriwritecloser
- /api/v2.7//uriwritecloser.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type URIWriteCloser

```go
type URIWriteCloser interface {
	io.WriteCloser

	URI() URI
}
```

URIWriteCloser represents a cross platform data writer for a file resource. This will normally refer to a local file resource.
