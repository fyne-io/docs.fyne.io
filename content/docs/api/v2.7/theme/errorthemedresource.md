---
tags: [api]
title: theme.ErrorThemedResource
slug: errorthemedresource

aliases:
- /api/v2.7/theme/errorthemedresource

package: fyne.io/fyne/v2/theme
---


---
```go
import "fyne.io/fyne/v2/theme"
```

#

###

```go
type ErrorThemedResource struct {
}
```

ErrorThemedResource is a resource wrapper that will return a version of the resource with the main color changed to indicate an error.

###

```go
func NewErrorThemedResource(orig fyne.Resource) *ErrorThemedResource
```
NewErrorThemedResource creates a resource that adapts to the error color for the current theme.

###

```go
func (res *ErrorThemedResource) Content() []byte
```
Content returns the underlying content of the resource adapted to the current background color.

###

```go
func (res *ErrorThemedResource) Name() string
```
Name returns the underlying resource name (used for caching).

###

```go
func (res *ErrorThemedResource) Original() fyne.Resource
```
Original returns the underlying resource that this error themed resource was adapted from

###

```go
func (res *ErrorThemedResource) ThemeColorName() fyne.ThemeColorName
```
ThemeColorName returns the fyne.ThemeColorName that is used as foreground color.


<div class="since">Since: <code>
2.6</code></div>
