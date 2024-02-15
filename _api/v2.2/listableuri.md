---
layout: page
tags: [api]
title: Fyne API "fyne.ListableURI"
---

# fyne.ListableURI
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ListableURI

```go
type ListableURI interface {
	URI

	// List returns a list of child URIs of this URI.
	List() ([]URI, error)
}
```

ListableURI represents a URI that can have child items, most commonly a directory on disk in the native filesystem.


<div class="since">Since: <code>
1.4</code></div>
