---
layout: page
tags: [api]
title: Fyne API storage
---

# storage
---
```go
import "fyne.io/fyne/storage"
```


## Usage

#### func  NewURI

```go
func NewURI(u string) fyne.URI
```
NewURI creates a new URI from the given string representation. This could be a URI from an external source or one saved from URI.String()

#### func  OpenFileFromURI

```go
func OpenFileFromURI(uri fyne.URI) (fyne.URIReadCloser, error)
```
OpenFileFromURI loads a file read stream from a resource identifier. This is mostly provided so that file references can be saved using their URI and loaded again later.

#### func  SaveFileToURI

```go
func SaveFileToURI(uri fyne.URI) (fyne.URIWriteCloser, error)
```
SaveFileToURI loads a file write stream to a resource identifier. This is mostly provided so that file references can be saved using their URI and written to again later.

#### types

 * [ExtensionFileFilter](extensionfilefilter.html)
 * [FileFilter](filefilter.html)
 * [MimeTypeFileFilter](mimetypefilefilter.html)
