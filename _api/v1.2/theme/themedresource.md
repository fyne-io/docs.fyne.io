---
layout: page
tags: [api]
title: Fyne API theme
---

# theme
---
```go
import "fyne.io/fyne/theme"
```

## Usage

#### type ThemedResource

```go
type ThemedResource struct {
}
```

ThemedResource is a resource wrapper that will return an appropriate resource for the currently selected theme.

#### func  NewThemedResource

```go
func NewThemedResource(src, ignored fyne.Resource) *ThemedResource
```
NewThemedResource creates a resource that adapts to the current theme setting.


<div class="deprecated">
Deprecated: NewThemedResource() will be replaced with a single parameter version in a future release</div> usage of this method will break, but using the first parameter only will be a trivial change.

#### func (*ThemedResource) Content

```go
func (res *ThemedResource) Content() []byte
```
Content returns the underlying content of the correct resource for the current theme

#### func (*ThemedResource) Name

```go
func (res *ThemedResource) Name() string
```
Name returns the underlying resource name (used for caching)
