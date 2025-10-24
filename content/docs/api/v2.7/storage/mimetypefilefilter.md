---
tags: [api]
title: storage.MimeTypeFileFilter
slug: mimetypefilefilter

aliases:
- /api/v2.7/storage/mimetypefilefilter

package: fyne.io/fyne/v2/storage
---


---
```go
import "fyne.io/fyne/v2/storage"
```

#

###

```go
type MimeTypeFileFilter struct {
	MimeTypes []string
}
```

MimeTypeFileFilter represents a file filter based on the files mime type, for example "image/*", "audio/mp3".

###

```go
func (mt *MimeTypeFileFilter) Matches(uri fyne.URI) bool
```
Matches returns true if a file URI has one of the filtered mimetypes.
