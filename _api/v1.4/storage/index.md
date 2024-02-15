---
layout: page
tags: [api]
title: Fyne API "storage"
---

# storage
---
```go
import "fyne.io/fyne/storage"
```

Package storage provides storage access and management functionality.

## Usage

```go
const URIRootError uriRootError = uriRootError("Cannot take the parent of the root element in a URI")
```
URIRootError should be thrown by fyne.URI implementations when the caller attempts to take the parent of the root. This way, downstream code that wants to programmatically walk up a URIs parent's will know when to stop iterating.

#### func  Child

```go
func Child(u fyne.URI, component string) (fyne.URI, error)
```
Child appends a new path element to a URI, separated by a '/' character.


<div class="since">Since: <code>
1.4</code></div>

#### func  Exists

```go
func Exists(u fyne.URI) (bool, error)
```
Exists will return true if the resource the URI refers to exists, and false otherwise. If an error occurs while checking, false is returned as the first return.


<div class="since">Since: <code>
1.4</code></div>

#### func  ListerForURI

```go
func ListerForURI(uri fyne.URI) (fyne.ListableURI, error)
```
ListerForURI will attempt to use the application's driver to convert a standard URI into a listable URI.


<div class="since">Since: <code>
1.4</code></div>

#### func  LoadResourceFromURI

```go
func LoadResourceFromURI(u fyne.URI) (fyne.Resource, error)
```
LoadResourceFromURI creates a new StaticResource in memory using the contents of the specified URI. The URI will be opened using the current driver, so valid schemas will vary from platform to platform. The file:// schema will always work.

#### func  NewFileURI

```go
func NewFileURI(path string) fyne.URI
```
NewFileURI creates a new URI from the given file path.

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

#### func  Parent

```go
func Parent(u fyne.URI) (fyne.URI, error)
```
Parent gets the parent of a URI by splitting it along '/' separators and removing the last item.


<div class="since">Since: <code>
1.4</code></div>

#### func  SaveFileToURI

```go
func SaveFileToURI(uri fyne.URI) (fyne.URIWriteCloser, error)
```
SaveFileToURI loads a file write stream to a resource identifier. This is mostly provided so that file references can be saved using their URI and written to again later.

#### types

 * [ExtensionFileFilter](extensionfilefilter.html)
 * [FileFilter](filefilter.html)
 * [MimeTypeFileFilter](mimetypefilefilter.html)
