---
layout: page
tags: [api]
title: Fyne API "fyne.URIReadCloser"
---

# fyne.URIReadCloser
---
```go
import "fyne.io/fyne"
```

## Usage

#### type URIReadCloser

```go
type URIReadCloser interface {
	io.ReadCloser
	// Deprecated, use URI().Name() instead
	Name() string
	URI() URI
}
```

URIReadCloser represents a cross platform data stream from a file or provider of data. It may refer to an item on a filesystem or data in another application that we have access to.
