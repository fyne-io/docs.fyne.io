---
tags: [api]
title: theme.DisabledResource
slug: disabledresource

aliases:
- /api/v2.7/theme/disabledresource

package: fyne.io/fyne/v2/theme
---


---
```go
import "fyne.io/fyne/v2/theme"
```

#

###

```go
type DisabledResource struct {
}
```

DisabledResource is a resource wrapper that will return an appropriate resource colorized by the current theme's `DisabledColor` color.

###

```go
func NewDisabledResource(res fyne.Resource) *DisabledResource
```
NewDisabledResource creates a resource that adapts to the current theme's DisabledColor setting.

###

```go
func (res *DisabledResource) Content() []byte
```
Content returns the disabled style content of the correct resource for the current theme

###

```go
func (res *DisabledResource) Name() string
```
Name returns the resource source name prefixed with `disabled_` (used for caching)

###

```go
func (res *DisabledResource) ThemeColorName() fyne.ThemeColorName
```
ThemeColorName returns the fyne.ThemeColorName that is used as foreground color.


<div class="since">Since: <code>
2.6</code></div>
