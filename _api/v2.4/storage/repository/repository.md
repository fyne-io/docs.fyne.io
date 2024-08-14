---
layout: page
tags: [api]
title: Fyne API "repository.Repository"
package: fyne.io/fyne/v2/storage/repository
---

# repository.Repository
---
```go
import "fyne.io/fyne/v2/storage/repository"
```

## Usage

#### type Repository

```go
type Repository interface {

	// Exists will be used to implement calls to storage.Exists() for the
	// registered scheme of this repository.
	//
	// Since: 2.0
	Exists(u fyne.URI) (bool, error)

	// Reader will be used to implement calls to storage.Reader()
	// for the registered scheme of this repository.
	//
	// Since: 2.0
	Reader(u fyne.URI) (fyne.URIReadCloser, error)

	// CanRead will be used to implement calls to storage.CanRead() for the
	// registered scheme of this repository.
	//
	// Since: 2.0
	CanRead(u fyne.URI) (bool, error)

	// Destroy is called when the repository is un-registered from a given
	// URI scheme.
	//
	// The string parameter will be the URI scheme that the repository was
	// registered for. This may be useful for repositories that need to
	// handle more than one URI scheme internally.
	//
	// Since: 2.0
	Destroy(string)
}
```

Repository represents a storage repository, which is a set of methods which implement specific functions on a URI. Repositories are registered to handle specific URI schemes, and the higher-level functions that operate on URIs internally look up an appropriate method from the relevant Repository.

The repository interface includes only methods which must be implemented at a minimum. Without implementing all of the methods in this interface, a URI would not be usable in a useful way. Additional functionality can be exposed by using interfaces which extend Repository.

Repositories are registered to handle a specific URI scheme (or schemes) using the Register() method. When a higher-level URI function such as storage.Copy() is called, the storage package will internally look up the repository associated with the scheme of the URI, then it will use a type assertion to check if the repository implements CopyableRepository. If so, the Copy() function will be run from the repository, otherwise storage.Copy() will return NotSupportedError. This works similarly for all other methods in repository-related interfaces.

Note that a repository can be registered for multiple URI schemes. In such cases, the repository must internally select and implement the correct behavior for each URI scheme.

A repository will only ever need to handle URIs with schemes for which it was registered, with the exception that functions with more than 1 operand such as Copy() and Move(), in which cases only the first operand is guaranteed to match a scheme for which the repository is registered.

NOTE: most developers who use Fyne should *not* generally attempt to call repository methods directly. You should use the methods in the storage package, which will automatically detect the scheme of a URI and call into the appropriate repository.


<div class="since">Since: <code>
2.0</code></div>

#### func  ForScheme

```go
func ForScheme(scheme string) (Repository, error)
```
ForScheme returns the Repository instance which is registered to handle URIs of the given scheme.

NOTE: this function is intended to be used specifically by the storage package. It generally should not be used outside of the fyne package - instead you should use the methods in the storage package.


<div class="since">Since: <code>
2.0</code></div>

#### func  ForURI

```go
func ForURI(u fyne.URI) (Repository, error)
```
ForURI returns the Repository instance which is registered to handle URIs of the given scheme. This is a helper method that calls ForScheme() on the scheme of the given URI.

NOTE: this function is intended to be used specifically by the storage package. It generally should not be used outside of the fyne package - instead you should use the methods in the storage package.


<div class="since">Since: <code>
2.0</code></div>
