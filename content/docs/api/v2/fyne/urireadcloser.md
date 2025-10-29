---
tags: [api]
title: fyne.URIReadCloser
slug: urireadcloser

aliases:
- /api//urireadcloser
- /api//urireadcloser.html
- /api/v2.0//urireadcloser
- /api/v2.0//urireadcloser.html
- /api/v2.1//urireadcloser
- /api/v2.1//urireadcloser.html
- /api/v2.2//urireadcloser
- /api/v2.2//urireadcloser.html
- /api/v2.3//urireadcloser
- /api/v2.3//urireadcloser.html
- /api/v2.4//urireadcloser
- /api/v2.4//urireadcloser.html
- /api/v2.5//urireadcloser
- /api/v2.5//urireadcloser.html
- /api/v2.6//urireadcloser
- /api/v2.6//urireadcloser.html
- /api/v2.7//urireadcloser
- /api/v2.7//urireadcloser.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type URIReadCloser

```go
type URIReadCloser interface {
	io.ReadCloser

	URI() URI
}
```

URIReadCloser represents a cross platform data stream from a file or provider of data. It may refer to an item on a filesystem or data in another application that we have access to.
