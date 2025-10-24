---
tags: [api]
title: theme.ThemedResource
slug: themedresource

aliases:
- /api/v2.0/theme/themedresource
- /api/v2.1/theme/themedresource
- /api/v2.2/theme/themedresource
- /api/v2.3/theme/themedresource
- /api/v2.4/theme/themedresource
- /api/v2.5/theme/themedresource
- /api/v2.6/theme/themedresource
- /api/v2.7/theme/themedresource

package: fyne.io/fyne/v2/theme
---


---
```go
import "fyne.io/fyne/v2/theme"
```

## Usage

#### type ThemedResource

```go
type ThemedResource struct {

	// ColorName specifies which theme colour should be used to theme the resource
	//
	// Since: 2.3
	ColorName fyne.ThemeColorName
}
```

ThemedResource is a resource wrapper that will return a version of the resource with the main color changed for the currently selected theme.

#### func  NewColoredResource

```go
func NewColoredResource(src fyne.Resource, name fyne.ThemeColorName) *ThemedResource
```
NewColoredResource creates a resource that adapts to the current theme setting using the color named in the constructor.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewSuccessThemedResource

```go
func NewSuccessThemedResource(src fyne.Resource) *ThemedResource
```
NewSuccessThemedResource creates a resource that adapts to the current theme success color.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewThemedResource

```go
func NewThemedResource(src fyne.Resource) *ThemedResource
```
NewThemedResource creates a resource that adapts to the current theme setting. By default, this will match the foreground color, but it can be changed using the `ColorName` field.

#### func  NewWarningThemedResource

```go
func NewWarningThemedResource(src fyne.Resource) *ThemedResource
```
NewWarningThemedResource creates a resource that adapts to the current theme warning color.


<div class="since">Since: <code>
2.4</code></div>

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

#### func (*ThemedResource) ThemeColorName

```go
func (res *ThemedResource) ThemeColorName() fyne.ThemeColorName
```
ThemeColorName returns the fyne.ThemeColorName that is used as foreground color.
