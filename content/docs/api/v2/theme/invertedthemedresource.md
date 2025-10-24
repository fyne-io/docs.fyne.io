---
tags: [api]
title: theme.InvertedThemedResource
slug: invertedthemedresource

aliases:
- /api/v2.0/theme/invertedthemedresource
- /api/v2.1/theme/invertedthemedresource
- /api/v2.2/theme/invertedthemedresource
- /api/v2.3/theme/invertedthemedresource
- /api/v2.4/theme/invertedthemedresource
- /api/v2.5/theme/invertedthemedresource
- /api/v2.6/theme/invertedthemedresource
- /api/v2.7/theme/invertedthemedresource

package: fyne.io/fyne/v2/theme
---


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

#### func (*InvertedThemedResource) ThemeColorName

```go
func (res *InvertedThemedResource) ThemeColorName() fyne.ThemeColorName
```
ThemeColorName returns the fyne.ThemeColorName that is used as foreground color.
