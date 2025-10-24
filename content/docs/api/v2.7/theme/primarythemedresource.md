---
tags: [api]
title: theme.PrimaryThemedResource
slug: primarythemedresource

aliases:
- /api/v2.7/theme/primarythemedresource

package: fyne.io/fyne/v2/theme
---


---
```go
import "fyne.io/fyne/v2/theme"
```

#

###

```go
type PrimaryThemedResource struct {
}
```

PrimaryThemedResource is a resource wrapper that will return a version of the resource with the main color changed to the theme primary color.

###

```go
func NewPrimaryThemedResource(orig fyne.Resource) *PrimaryThemedResource
```
NewPrimaryThemedResource creates a resource that adapts to the primary color for the current theme.

###

```go
func (res *PrimaryThemedResource) Content() []byte
```
Content returns the underlying content of the resource adapted to the current background color.

###

```go
func (res *PrimaryThemedResource) Name() string
```
Name returns the underlying resource name (used for caching).

###

```go
func (res *PrimaryThemedResource) Original() fyne.Resource
```
Original returns the underlying resource that this primary themed resource was adapted from

###

```go
func (res *PrimaryThemedResource) ThemeColorName() fyne.ThemeColorName
```
ThemeColorName returns the fyne.ThemeColorName that is used as foreground color.


<div class="since">Since: <code>
2.6</code></div>
