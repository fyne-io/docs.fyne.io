---
tags: [api]
title: storage.ExtensionFileFilter
slug: extensionfilefilter

aliases:
- /api/v2.7/storage/extensionfilefilter

package: fyne.io/fyne/v2/storage
---


---
```go
import "fyne.io/fyne/v2/storage"
```

#

###

```go
type ExtensionFileFilter struct {
	Extensions []string
}
```

ExtensionFileFilter represents a file filter based on the ending of file names, for example ".txt" and ".png".

###

```go
func (e *ExtensionFileFilter) Matches(uri fyne.URI) bool
```
Matches returns true if a file URI has one of the filtered extensions.
