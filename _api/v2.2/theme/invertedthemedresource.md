---
layout: page
tags: [api]
title: Fyne API "theme.InvertedThemedResource"
---

# theme.InvertedThemedResource
---
```go
import "fyne.io/fyne/v2/theme"
```

## Usage

#### type InvertedThemedResource

```go
type InvertedThemedResource struct {
}
```

InvertedThemedResource is a resource wrapper that will return a version of the resource with the main color changed for use over highlighted elements.

#### func  NewInvertedThemedResource

```go
func NewInvertedThemedResource(orig fyne.Resource) *InvertedThemedResource
```
NewInvertedThemedResource creates a resource that adapts to the current theme for use over highlighted elements.

#### func (*InvertedThemedResource) Content

```go
func (res *InvertedThemedResource) Content() []byte
```
Content returns the underlying content of the resource adapted to the current background color.

#### func (*InvertedThemedResource) Name

```go
func (res *InvertedThemedResource) Name() string
```
Name returns the underlying resource name (used for caching).

#### func (*InvertedThemedResource) Original

```go
func (res *InvertedThemedResource) Original() fyne.Resource
```
Original returns the underlying resource that this inverted themed resource was adapted from
