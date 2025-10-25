---
tags: [api]
title: storage.FileFilter
slug: filefilter

aliases:
- /api/v2/storage/filefilter.html
- /api/v2.0/storage/filefilter
- /api/v2.0/storage/filefilter.html
- /api/v2.1/storage/filefilter
- /api/v2.1/storage/filefilter.html
- /api/v2.2/storage/filefilter
- /api/v2.2/storage/filefilter.html
- /api/v2.3/storage/filefilter
- /api/v2.3/storage/filefilter.html
- /api/v2.4/storage/filefilter
- /api/v2.4/storage/filefilter.html
- /api/v2.5/storage/filefilter
- /api/v2.5/storage/filefilter.html
- /api/v2.6/storage/filefilter
- /api/v2.6/storage/filefilter.html
- /api/v2.7/storage/filefilter
- /api/v2.7/storage/filefilter.html

package: fyne.io/fyne/v2/storage
---


---
```go
import "fyne.io/fyne/v2/storage"
```

## Usage

#### type FileFilter

```go
type FileFilter interface {
	Matches(fyne.URI) bool
}
```

FileFilter is an interface that can be implemented to provide a filter to a file dialog.

#### func  NewExtensionFileFilter

```go
func NewExtensionFileFilter(extensions []string) FileFilter
```
NewExtensionFileFilter takes a string slice of extensions with a leading . and creates a filter for the file dialog. Example: .jpg, .mp3, .txt, .sh

#### func  NewMimeTypeFileFilter

```go
func NewMimeTypeFileFilter(mimeTypes []string) FileFilter
```
NewMimeTypeFileFilter takes a string slice of mimetypes, including globs, and creates a filter for the file dialog. Example: image/*, audio/mp3, text/plain, application/*
