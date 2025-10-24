---
tags: [api]
title: fyne.ListableURI
slug: listableuri

aliases:
- /api/v2.7//listableuri

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

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
