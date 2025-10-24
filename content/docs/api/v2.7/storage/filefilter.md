---
tags: [api]
title: storage.FileFilter
slug: filefilter

aliases:
- /api/v2.7/storage/filefilter

package: fyne.io/fyne/v2/storage
---


---
```go
import "fyne.io/fyne/v2/storage"
```

#

###

```go
type FileFilter interface {
	Matches(fyne.URI) bool
}
```

FileFilter is an interface that can be implemented to provide a filter to a file dialog.

###

```go
func NewExtensionFileFilter(extensions []string) FileFilter
```
NewExtensionFileFilter takes a string slice of extensions with a leading . and creates a filter for the file dialog. Example: .jpg, .mp3, .txt, .sh

###

```go
func NewMimeTypeFileFilter(mimeTypes []string) FileFilter
```
NewMimeTypeFileFilter takes a string slice of mimetypes, including globs, and creates a filter for the file dialog. Example: image/*, audio/mp3, text/plain, application/*
