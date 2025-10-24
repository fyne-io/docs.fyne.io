---
tags: [api]
title: storage.MimeTypeFileFilter
slug: mimetypefilefilter

aliases:
- /api/v2.0/storage/mimetypefilefilter
- /api/v2.1/storage/mimetypefilefilter
- /api/v2.2/storage/mimetypefilefilter
- /api/v2.3/storage/mimetypefilefilter
- /api/v2.4/storage/mimetypefilefilter
- /api/v2.5/storage/mimetypefilefilter
- /api/v2.6/storage/mimetypefilefilter
- /api/v2.7/storage/mimetypefilefilter

package: fyne.io/fyne/v2/storage
---


---
```go
import "fyne.io/fyne/v2/storage"
```

## Usage

#### type MimeTypeFileFilter

```go
type MimeTypeFileFilter struct {
	MimeTypes []string
}
```

MimeTypeFileFilter represents a file filter based on the files mime type, for example "image/*", "audio/mp3".

#### func (*MimeTypeFileFilter) Matches

```go
func (mt *MimeTypeFileFilter) Matches(uri fyne.URI) bool
```
Matches returns true if a file URI has one of the filtered mimetypes.
