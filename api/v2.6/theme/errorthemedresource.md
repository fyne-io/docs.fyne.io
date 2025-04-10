---
layout: page
tags: [api]
title: Fyne API "theme.ErrorThemedResource"
package: fyne.io/fyne/v2/theme
---

# theme.ErrorThemedResource
---
```go
import "fyne.io/fyne/v2/theme"
```

## Usage

#### type ErrorThemedResource

```go
type ErrorThemedResource struct {
}
```

ErrorThemedResource is a resource wrapper that will return a version of the resource with the main color changed to indicate an error.

#### func  NewErrorThemedResource

```go
func NewErrorThemedResource(orig fyne.Resource) *ErrorThemedResource
```
NewErrorThemedResource creates a resource that adapts to the error color for the current theme.

#### func (*ErrorThemedResource) Content

```go
func (res *ErrorThemedResource) Content() []byte
```
Content returns the underlying content of the resource adapted to the current background color.

#### func (*ErrorThemedResource) Name

```go
func (res *ErrorThemedResource) Name() string
```
Name returns the underlying resource name (used for caching).

#### func (*ErrorThemedResource) Original

```go
func (res *ErrorThemedResource) Original() fyne.Resource
```
Original returns the underlying resource that this error themed resource was adapted from

#### func (*ErrorThemedResource) ThemeColorName

```go
func (res *ErrorThemedResource) ThemeColorName() fyne.ThemeColorName
```
ThemeColorName returns the fyne.ThemeColorName that is used as foreground color. @implements fyne.ThemedResource


<div class="since">Since: <code>
2.6</code></div>
