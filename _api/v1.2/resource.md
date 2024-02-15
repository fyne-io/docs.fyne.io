---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
---
```go
import "fyne.io/fyne"
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
LoadResourceFromPath creates a new StaticResource in memory using the contents of the specified file.

#### func  LoadResourceFromURLString

```go
func LoadResourceFromURLString(urlStr string) (Resource, error)
```
LoadResourceFromURLString creates a new StaticResource in memory using the body of the specified URL.
