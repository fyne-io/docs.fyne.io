---
layout: page
tags: [api]
title: Fyne API "repository"
---

# repository
---
```go
import "fyne.io/fyne/v2/storage/repository"
```


## Usage

```go
var (
	// ErrOperationNotSupported may be thrown by certain functions in the storage
	// or repository packages which operate on URIs if an operation is attempted
	// that is not supported for the scheme relevant to the URI, normally because
	// the underlying repository has either not implemented the relevant function,
	// or has explicitly returned this error.
	//
	// Since: 2.0
	ErrOperationNotSupported = errors.New("operation not supported for this URI")

	// ErrURIRoot should be thrown by fyne.URI implementations when the caller
	// attempts to take the parent of the root. This way, downstream code that
	// wants to programmatically walk up a URIs parent's will know when to stop
	// iterating.
	//
	// Since: 2.0
	ErrURIRoot = errors.New("cannot take the parent of the root element in a URI")
)
```

#### func  GenericChild

```go
func GenericChild(u fyne.URI, component string) (fyne.URI, error)
```
GenericChild can be used as a common-case implementation of HierarchicalRepository.Child(). It will create a child URI by separating the URI into it's component parts as described in IETF RFC 3986, then appending "/" + component to the path, then concatenating the result and parsing it as a new URI.

NOTE: this function should not be called except by an implementation of the Repository interface - using this for unknown URIs may break.


<div class="since">Since: <code>
2.0</code></div>

#### func  GenericCopy

```go
func GenericCopy(source fyne.URI, destination fyne.URI) error
```
GenericCopy can be used a common-case implementation of CopyableRepository.Copy(). It will perform the copy by obtaining a reader for the source URI, a writer for the destination URI, then writing the contents of the source to the destination.

For obvious reasons, the destination URI must have a registered WritableRepository.

NOTE: this function should not be called except by an implementation of the Repository interface - using this for unknown URIs may break.


<div class="since">Since: <code>
2.0</code></div>

#### func  GenericMove

```go
func GenericMove(source fyne.URI, destination fyne.URI) error
```
GenericMove can be used a common-case implementation of MovableRepository.Move(). It will perform the move by obtaining a reader for the source URI, a writer for the destination URI, then writing the contents of the source to the destination. Following this, the source will be deleted using WritableRepository.Delete.

For obvious reasons, the source and destination URIs must both be writable.

NOTE: this function should not be called except by an implementation of the Repository interface - using this for unknown URIs may break.


<div class="since">Since: <code>
2.0</code></div>

#### func  GenericParent

```go
func GenericParent(u fyne.URI) (fyne.URI, error)
```
GenericParent can be used as a common-case implementation of HierarchicalRepository.Parent(). It will create a parent URI based on IETF RFC3986.

In short, the URI is separated into it's component parts, the path component is split along instances of '/', and the trailing element is removed. The result is concatenated and parsed as a new URI.

If the URI path is empty or '/', then a nil URI is returned, along with ErrURIRoot.

NOTE: this function should not be called except by an implementation of the Repository interface - using this for unknown URIs may break.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewFileURI

```go
func NewFileURI(path string) fyne.URI
```
NewFileURI implements the back-end logic to storage.NewFileURI, which you should use instead. This is only here because other functions in repository need to call it, and it prevents a circular import.


<div class="since">Since: <code>
2.0</code></div>

#### func  ParseURI

```go
func ParseURI(s string) (fyne.URI, error)
```
ParseURI implements the back-end logic for storage.ParseURI, which you should use instead. This is only here because other functions in repository need to call it, and it prevents a circular import.


<div class="since">Since: <code>
2.0</code></div>

#### func  Register

```go
func Register(scheme string, repository Repository)
```
Register registers a storage repository so that operations on URIs of the registered scheme will use methods implemented by the relevant repository implementation.


<div class="since">Since: <code>
2.0</code></div>

#### types

 * [CopyableRepository](copyablerepository.html)
 * [CustomURIRepository](customurirepository.html)
 * [HierarchicalRepository](hierarchicalrepository.html)
 * [ListableRepository](listablerepository.html)
 * [MovableRepository](movablerepository.html)
 * [Repository](repository.html)
 * [WritableRepository](writablerepository.html)
