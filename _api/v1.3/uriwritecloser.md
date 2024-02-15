---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
---
```go
import "fyne.io/fyne"
```

## Usage

#### type URIWriteCloser

```go
type URIWriteCloser interface {
	io.WriteCloser
	Name() string
	URI() URI
}
```

URIWriteCloser represents a cross platform data writer for a file resource. This will normally refer to a local file resource.
