---
layout: page
tags: [api]
title: Fyne API "theme.PrimaryThemedResource"
---

# theme.PrimaryThemedResource
---
```go
import "fyne.io/fyne/v2/theme"
```

## Usage

#### type PrimaryThemedResource

```go
type PrimaryThemedResource struct {
}
```

PrimaryThemedResource is a resource wrapper that will return a version of the resource with the main color changed to the theme primary color.

#### func  NewPrimaryThemedResource

```go
func NewPrimaryThemedResource(orig fyne.Resource) *PrimaryThemedResource
```
NewPrimaryThemedResource creates a resource that adapts to the primary color for the current theme.

#### func (*PrimaryThemedResource) Content

```go
func (res *PrimaryThemedResource) Content() []byte
```
Content returns the underlying content of the resource adapted to the current background color.

#### func (*PrimaryThemedResource) Name

```go
func (res *PrimaryThemedResource) Name() string
```
Name returns the underlying resource name (used for caching).

#### func (*PrimaryThemedResource) Original

```go
func (res *PrimaryThemedResource) Original() fyne.Resource
```
Original returns the underlying resource that this primary themed resource was adapted from
