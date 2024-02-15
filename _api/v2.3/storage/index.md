---
layout: page
tags: [api]
title: Fyne API "storage"
---

# storage
---
```go
import "fyne.io/fyne/v2/storage"
```

Package storage provides storage access and management functionality.

## Usage

```go
var (
	// ErrAlreadyExists may be thrown by docs. E.g., save a document twice.
	//
	// Since: 2.3
	ErrAlreadyExists = errors.New("document already exists")

	// ErrNotExists may be thrown by docs. E.g., save an unknown document.
	//
	// Since: 2.3
	ErrNotExists = errors.New("document does not exist")
)
```

```go
var URIRootError = repository.ErrURIRoot
```
URIRootError is a wrapper for repository.URIRootError

Deprecated - use repository.ErrURIRoot instead

#### func  CanList

```go
func CanList(u fyne.URI) (bool, error)
```
CanList will determine if the URI is listable or not.

This method may fail in several ways:

```go
    - Different permissions or credentials are required to check if the
      URI supports listing.

    - This URI scheme could represent some resources that can be listed,
      but this specific URI is not one of them (e.g. a file on a
      filesystem, as opposed to a directory).

    - Checking for listability depended on a lower level operation
      such as network or filesystem access that has failed in some way.

    - If the scheme of the given URI does not have a registered
      ListableRepository instance, then this method will fail with a
      repository.ErrOperationNotSupported.
```

CanList is backed by the repository system - this function calls into a scheme-specific implementation from a registered repository.


<div class="since">Since: <code>
2.0</code></div>

#### func  CanRead

```go
func CanRead(u fyne.URI) (bool, error)
```
CanRead determines if a given URI could be written to using the Reader() method. It is preferred to check if a URI is readable using this method before calling Reader(), because the underlying operations required to attempt to read and then report an error may be slower than the operations needed to test if a URI is readable. Keep in mind however that even if CanRead returns true, you must still do appropriate error handling for Reader(), as the underlying filesystem may have changed since you called CanRead.

The non-existence of a resource should not be treated as an error. In other words, a Repository implementation which for some URI u returns false, nil for Exists(u), CanRead(u) should also return false, nil.

CanRead is backed by the repository system - this function calls into a scheme-specific implementation from a registered repository.


<div class="since">Since: <code>
2.0</code></div>

#### func  CanWrite

```go
func CanWrite(u fyne.URI) (bool, error)
```
CanWrite is backed by the repository system - this function calls into a scheme-specific implementation from a registered repository.


<div class="since">Since: <code>
2.0</code></div>

#### func  Child

```go
func Child(u fyne.URI, component string) (fyne.URI, error)
```
Child returns a URI referencing a resource nested hierarchically below the given URI, identified by a string. For example, the child with the string component 'quux' of 'file://foo/bar' is 'file://foo/bar/quux'.

This can fail in several ways:

```go
    - If the URI refers to a resource which does not exist in a hierarchical
      context (e.g. the URI references something which does not have a
      semantically meaningful "child"), the Child() implementation may return an
      error.

    - If generating a reference to a child of the referenced resource requires
      interfacing with some external system, failures may propagate through the
      Child() implementation. It is expected that this case would occur very
      rarely if ever.

    - If the scheme of the given URI does not have a registered
      HierarchicalRepository instance, then this method will fail with a
      repository.ErrOperationNotSupported.
```

NOTE: since v2.0.0, Child() is backed by the repository system - this function is a helper which calls into an appropriate repository instance for the scheme of the URI it is given.


<div class="since">Since: <code>
1.4</code></div>

#### func  Copy

```go
func Copy(source fyne.URI, destination fyne.URI) error
```
Copy given two URIs, 'src', and 'dest' both of the same scheme, will copy one to the other. If the source and destination are of different schemes, then the Copy implementation for the storage repository registered to the scheme of the source will be used. Implementations are recommended to use repository.GenericCopy() as a fail-over in the case that they do not understand how to operate on the scheme of the destination URI. However, the behavior of calling Copy() on URIs of non-matching schemes is ultimately defined by the storage repository registered to the scheme of the source URI.

This method may fail in several ways:

```go
    - Different permissions or credentials are required to perform the
      copy operation.

    - This URI scheme could represent some resources that can be copied,
      but either the source, destination, or both are not resources
      that support copying.

    - Performing the copy operation depended on a lower level operation
      such as network or filesystem access that has failed in some way.

    - If the scheme of the given URI does not have a registered
      CopyableRepository instance, then this method will fail with a
      repository.ErrOperationNotSupported.
```

Copy is backed by the repository system - this function calls into a scheme-specific implementation from a registered repository.


<div class="since">Since: <code>
2.0</code></div>

#### func  CreateListable

```go
func CreateListable(u fyne.URI) error
```
CreateListable creates a new listable resource referenced by the given URI. CreateListable will error if the URI already references an extant resource. This method is used for storage repositories where listable resources are of a different underlying type than other resources - for example, in a typical filesystem ('file://'), CreateListable() corresponds to directory creation, and Writer() implies file creation for non-extant operands.

For storage repositories where listable and non-listable resources are the of the same underlying type, CreateListable should be equivalent to calling Writer(), writing zero bytes, and then closing the `URIWriteCloser - in filesystem terms, the same as calling 'touch;'.

Storage repositories which support listing, but not creation of listable objects may return repository.ErrOperationNotSupported.

CreateListable should generally fail if the parent of it's operand does not exist, however this can vary by the implementation details of the specific storage repository. In filesystem terms, this function is "mkdir" not "mkdir -p".

This method may fail in several ways:

```go
    - Different permissions or credentials are required to create the requested
      resource.

    - Creating the resource depended on a lower level operation such as network
      or filesystem access that has failed in some way.

    - If the scheme of the given URI does not have a registered
      ListableRepository instance, then this method will fail with a
      repository.ErrOperationNotSupported.
```

CreateListable is backed by the repository system - this function either calls into a scheme-specific implementation from a registered repository, or fails with a URIOperationNotSupported error.


<div class="since">Since: <code>
2.0</code></div>

#### func  Delete

```go
func Delete(u fyne.URI) error
```
Delete destroys, deletes, or otherwise removes the resource referenced by the URI.

This can fail in several ways:

```go
    - If removing the resource requires interfacing with some external system,
      failures may propagate through Destroy(). For example, deleting a file may
      fail with a permissions error.

    - If the referenced resource does not exist, attempting to destroy it should
      throw an error.

    - If the scheme of the given URI does not have a registered
      WritableRepository instance, then this method will fail with a
      repository.ErrOperationNotSupported.
```

Delete is backed by the repository system - this function calls into a scheme-specific implementation from a registered repository.


<div class="since">Since: <code>
2.0</code></div>

#### func  Exists

```go
func Exists(u fyne.URI) (bool, error)
```
Exists determines if the resource referenced by the URI exists.

This can fail in several ways:

```go
    - If checking the existence of a resource requires interfacing with some
      external system, then failures may propagate through Exists(). For
      example, checking the existence of a resource requires reading a directory
      may result in a permissions error.
```

It is understood that a non-nil error value signals that the existence or non-existence of the resource cannot be determined and is undefined.

NOTE: since v2.0.0, Exists is backed by the repository system - this function calls into a scheme-specific implementation from a registered repository.

Exists may call into either a generic implementation, or into a scheme-specific implementation depending on which storage repositories have been registered.


<div class="since">Since: <code>
1.4</code></div>

#### func  List

```go
func List(u fyne.URI) ([]fyne.URI, error)
```
List returns a list of URIs that reference resources which are nested below the resource referenced by the argument. For example, listing a directory on a filesystem should return a list of files and directories it contains.

This method may fail in several ways:

```go
    - Different permissions or credentials are required to obtain a
      listing for the given URI.

    - This URI scheme could represent some resources that can be listed,
      but this specific URI is not one of them (e.g. a file on a
      filesystem, as opposed to a directory). This can be tested in advance
      using the Listable() function.

    - Obtaining the listing depended on a lower level operation such as
      network or filesystem access that has failed in some way.

    - If the scheme of the given URI does not have a registered
      ListableRepository instance, then this method will fail with a
      repository.ErrOperationNotSupported.
```

List is backed by the repository system - this function either calls into a scheme-specific implementation from a registered repository, or fails with a URIOperationNotSupported error.


<div class="since">Since: <code>
2.0</code></div>

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

#### func  Move

```go
func Move(source fyne.URI, destination fyne.URI) error
```
Move returns a method that given two URIs, 'src' and 'dest' both of the same scheme this will move src to dest. This means the resource referenced by src will be copied into the resource referenced by dest, and the resource referenced by src will no longer exist after the operation is complete.

If the source and destination are of different schemes, then the Move implementation for the storage repository registered to the scheme of the source will be used. Implementations are recommended to use repository.GenericMove() as a fail-over in the case that they do not understand how to operate on the scheme of the destination URI. However, the behavior of calling Move() on URIs of non-matching schemes is ultimately defined by the storage repository registered to the scheme of the source URI.

This method may fail in several ways:

```go
    - Different permissions or credentials are required to perform the
      rename operation.

    - This URI scheme could represent some resources that can be renamed,
      but either the source, destination, or both are not resources
      that support renaming.

    - Performing the rename operation depended on a lower level operation
      such as network or filesystem access that has failed in some way.

    - If the scheme of the given URI does not have a registered
      MovableRepository instance, then this method will fail with a
      repository.ErrOperationNotSupported.
```

Move is backed by the repository system - this function calls into a scheme-specific implementation from a registered repository.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewFileURI

```go
func NewFileURI(path string) fyne.URI
```
NewFileURI creates a new URI from the given file path.

#### func  NewURI

```go
func NewURI(s string) fyne.URI
```
NewURI creates a new URI from the given string representation. This could be a URI from an external source or one saved from URI.String()


<div class="deprecated">
Deprecated: use ParseURI instead</div>

#### func  OpenFileFromURI

```go
func OpenFileFromURI(uri fyne.URI) (fyne.URIReadCloser, error)
```
OpenFileFromURI loads a file read stream from a resource identifier. This is mostly provided so that file references can be saved using their URI and loaded again later.


<div class="deprecated">
Deprecated: this has been replaced by storage.Reader(URI)</div>

#### func  Parent

```go
func Parent(u fyne.URI) (fyne.URI, error)
```
Parent returns a URI referencing the parent resource of the resource referenced by the URI. For example, the Parent() of 'file://foo/bar.baz' is 'file://foo'. The URI which is returned will be listable.

NOTE: it is not a given that Parent() return a parent URI with the same Scheme(), though this will normally be the case.

This can fail in several ways:

```go
    - If the URI refers to a filesystem root, then the Parent() implementation
      must return (nil, URIRootError).

    - If the URI refers to a resource which does not exist in a hierarchical
      context (e.g. the URI references something which does not have a
      semantically meaningful "parent"), the Parent() implementation may return
      an error.

    - If determining the parent of the referenced resource requires
      interfacing with some external system, failures may propagate
      through the Parent() implementation. For example if determining
      the parent of a file:// URI requires reading information from
      the filesystem, it could fail with a permission error.

    - If the scheme of the given URI does not have a registered
      HierarchicalRepository instance, then this method will fail with a
      repository.ErrOperationNotSupported.
```

NOTE: since v2.0.0, Parent() is backed by the repository system - this function is a helper which calls into an appropriate repository instance for the scheme of the URI it is given.


<div class="since">Since: <code>
1.4</code></div>

#### func  ParseURI

```go
func ParseURI(s string) (fyne.URI, error)
```
ParseURI creates a new URI instance by parsing a URI string.

Parse URI will parse up to the first ':' present in the URI string to extract the scheme, and then delegate further parsing to the registered repository for the given scheme. If no repository is registered for that scheme, the URI is parsed on a best-effort basis using net/url.

As a special exception, URIs beginning with 'file:' are always parsed using NewFileURI(), which will correctly handle back-slashes appearing in the URI path component on Windows.


<div class="since">Since: <code>
2.0</code></div>

#### func  Reader

```go
func Reader(u fyne.URI) (fyne.URIReadCloser, error)
```
Reader returns URIReadCloser set up to read from the resource that the URI references.

This method can fail in several ways:

```go
    - Different permissions or credentials are required to read the
      referenced resource.

    - This URI scheme could represent some resources that can be read,
      but this particular URI references a resources that is not
      something that can be read.

    - Attempting to set up the reader depended on a lower level
      operation such as a network or filesystem access that has failed
      in some way.
```

Reader is backed by the repository system - this function calls into a scheme-specific implementation from a registered repository.


<div class="since">Since: <code>
2.0</code></div>

#### func  SaveFileToURI

```go
func SaveFileToURI(uri fyne.URI) (fyne.URIWriteCloser, error)
```
SaveFileToURI loads a file write stream to a resource identifier. This is mostly provided so that file references can be saved using their URI and written to again later.


<div class="deprecated">
Deprecated: this has been replaced by storage.Writer(URI)</div>

#### func  Writer

```go
func Writer(u fyne.URI) (fyne.URIWriteCloser, error)
```
Writer returns URIWriteCloser set up to write to the resource that the URI references.

Writing to a non-extant resource should create that resource if possible (and if not possible, this should be reflected in the return of CanWrite()). Writing to an extant resource should overwrite it in-place. At present, this API does not provide a mechanism for appending to an already-extant resource, except for reading it in and writing all the data back out.

This method can fail in several ways:

```go
    - Different permissions or credentials are required to write to the
      referenced resource.

    - This URI scheme could represent some resources that can be
      written, but this particular URI references a resources that is
      not something that can be written.

    - Attempting to set up the writer depended on a lower level
      operation such as a network or filesystem access that has failed
      in some way.

    - If the scheme of the given URI does not have a registered
      WritableRepository instance, then this method will fail with a
      repository.ErrOperationNotSupported.
```

Writer is backed by the repository system - this function calls into a scheme-specific implementation from a registered repository.


<div class="since">Since: <code>
2.0</code></div>

#### types

 * [ExtensionFileFilter](extensionfilefilter.html)
 * [FileFilter](filefilter.html)
 * [MimeTypeFileFilter](mimetypefilefilter.html)
