---
layout: page
tags: [api]
title: Fyne API "theme.ThemedResource"
---

# theme.ThemedResource
---
```go
import "fyne.io/fyne/v2/theme"
```

## Usage

#### type ThemedResource

```go
type ThemedResource struct {
}
```

ThemedResource is a resource wrapper that will return a version of the resource with the main color changed for the currently selected theme.

#### func  NewThemedResource

```go
func NewThemedResource(src fyne.Resource) *ThemedResource
```
NewThemedResource creates a resource that adapts to the current theme setting.

#### func (*ThemedResource) Content

```go
func (res *ThemedResource) Content() []byte
```
Content returns the underlying content of the resource adapted to the current text color.

#### func (*ThemedResource) Error

```go
func (res *ThemedResource) Error() *ErrorThemedResource
```
Error returns a different resource for indicating an error.

#### func (*ThemedResource) Name

```go
func (res *ThemedResource) Name() string
```
Name returns the underlying resource name (used for caching).
