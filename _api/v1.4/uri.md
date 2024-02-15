---
layout: page
tags: [api]
title: Fyne API "fyne.URI"
---

# fyne.URI
---
```go
import "fyne.io/fyne"
```

## Usage

#### type URI

```go
type URI interface {
	fmt.Stringer
	Extension() string
	Name() string
	MimeType() string
	Scheme() string
}
```

URI represents the identifier of a resource on a target system. This resource may be a file or another data source such as an app or file sharing system.
