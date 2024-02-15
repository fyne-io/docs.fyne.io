---
layout: page
tags: [api]
title: Fyne API "fyne.URI"
---

# fyne.URI
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type URI

```go
type URI interface {
	fmt.Stringer

	// Extension should return the file extension of the resource
	// referenced by the URI. For example, the Extension() of
	// 'file://foo/bar.baz' is 'baz'. May return an empty string if the
	// referenced resource has none.
	Extension() string

	// Name should return the base name of the item referenced by the URI.
	// For example, the Name() of 'file://foo/bar.baz' is 'bar.baz'.
	Name() string

	// MimeType should return the content type of the resource referenced
	// by the URI. The returned string should be in the format described
	// by Section 5 of RFC2045 ("Content-Type Header Field").
	MimeType() string

	// Scheme should return the URI scheme of the URI as defined by IETF
	// RFC3986. For example, the Scheme() of 'file://foo/bar.baz` is
	// 'file'.
	//
	// Scheme should always return the scheme in all lower-case characters.
	Scheme() string

	// Authority should return the URI authority, as defined by IETF
	// RFC3986.
	//
	// NOTE: the RFC3986 can be obtained by combining the User and Host
	// Fields of net/url's URL structure. Consult IETF RFC3986, section
	// 3.2, pp. 17.
	//
	// Since: 2.0
	Authority() string

	// Path should return the URI path, as defined by IETF RFC3986.
	//
	// Since: 2.0
	Path() string

	// Query should return the URI query, as defined by IETF RFC3986.
	//
	// Since: 2.0
	Query() string

	// Fragment should return the URI fragment, as defined by IETF
	// RFC3986.
	//
	// Since: 2.0
	Fragment() string
}
```

URI represents the identifier of a resource on a target system. This resource may be a file or another data source such as an app or file sharing system.

In general, it is expected that URI implementations follow IETF RFC3896. Implementations are highly recommended to utilize net/url to implement URI parsing methods, especially Scheme(), AUthority(), Path(), Query(), and Fragment().
