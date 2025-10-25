---
tags: [api]
title: fyne.ListableURI
slug: listableuri

aliases:
- /api/v2/fyne/listableuri.html
- /api/v2.0//listableuri
- /api/v2.0//listableuri.html
- /api/v2.1//listableuri
- /api/v2.1//listableuri.html
- /api/v2.2//listableuri
- /api/v2.2//listableuri.html
- /api/v2.3//listableuri
- /api/v2.3//listableuri.html
- /api/v2.4//listableuri
- /api/v2.4//listableuri.html
- /api/v2.5//listableuri
- /api/v2.5//listableuri.html
- /api/v2.6//listableuri
- /api/v2.6//listableuri.html
- /api/v2.7//listableuri
- /api/v2.7//listableuri.html

package: fyne.io/fyne/v2
---


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

ListableURI represents a [URI] that can have child items, most commonly a directory on disk in the native filesystem.


<div class="since">Since: <code>
1.4</code></div>
