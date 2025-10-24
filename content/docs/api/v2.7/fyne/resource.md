---
tags: [api]
title: fyne.Resource
slug: resource

aliases:
- /api/v2.7//resource

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Resource interface {
	Name() string
	Content() []byte
}
```

Resource represents a single binary resource, such as an image or font. A resource has an identifying name and byte array content. The serialised path of a resource can be obtained which may result in a blocking filesystem write operation.

###

```go
func LoadResourceFromPath(path string) (Resource, error)
```
LoadResourceFromPath creates a new [StaticResource] in memory using the contents of the specified file.

###

```go
func LoadResourceFromURLString(urlStr string) (Resource, error)
```
LoadResourceFromURLString creates a new [StaticResource] in memory using the body of the specified URL.
