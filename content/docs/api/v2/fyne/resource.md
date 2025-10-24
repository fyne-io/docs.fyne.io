---
tags: [api]
title: fyne.Resource
slug: resource

aliases:
- /api/v2.0//resource
- /api/v2.1//resource
- /api/v2.2//resource
- /api/v2.3//resource
- /api/v2.4//resource
- /api/v2.5//resource
- /api/v2.6//resource
- /api/v2.7//resource

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Resource

```go
type Resource interface {
	Name() string
	Content() []byte
}
```

Resource represents a single binary resource, such as an image or font. A resource has an identifying name and byte array content. The serialised path of a resource can be obtained which may result in a blocking filesystem write operation.

#### func  LoadResourceFromPath

```go
func LoadResourceFromPath(path string) (Resource, error)
```
LoadResourceFromPath creates a new [StaticResource] in memory using the contents of the specified file.

#### func  LoadResourceFromURLString

```go
func LoadResourceFromURLString(urlStr string) (Resource, error)
```
LoadResourceFromURLString creates a new [StaticResource] in memory using the body of the specified URL.
