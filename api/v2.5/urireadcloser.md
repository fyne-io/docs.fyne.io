---
layout: page
tags: [api]
title: Fyne API "fyne.URIReadCloser"
package: fyne.io/fyne/v2
---

# fyne.URIReadCloser
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
