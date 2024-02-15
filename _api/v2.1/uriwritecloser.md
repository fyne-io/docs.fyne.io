---
layout: page
tags: [api]
title: Fyne API "fyne.URIWriteCloser"
---

# fyne.URIWriteCloser
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
