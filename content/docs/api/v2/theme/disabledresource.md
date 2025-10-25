---
tags: [api]
title: theme.DisabledResource
slug: disabledresource

aliases:
- /api/v2/theme/disabledresource.html
- /api/v2.0/theme/disabledresource
- /api/v2.0/theme/disabledresource.html
- /api/v2.1/theme/disabledresource
- /api/v2.1/theme/disabledresource.html
- /api/v2.2/theme/disabledresource
- /api/v2.2/theme/disabledresource.html
- /api/v2.3/theme/disabledresource
- /api/v2.3/theme/disabledresource.html
- /api/v2.4/theme/disabledresource
- /api/v2.4/theme/disabledresource.html
- /api/v2.5/theme/disabledresource
- /api/v2.5/theme/disabledresource.html
- /api/v2.6/theme/disabledresource
- /api/v2.6/theme/disabledresource.html
- /api/v2.7/theme/disabledresource
- /api/v2.7/theme/disabledresource.html

package: fyne.io/fyne/v2/theme
---


---
```go
import "fyne.io/fyne/v2/theme"
```

## Usage

#### type DisabledResource

```go
type DisabledResource struct {
}
```

DisabledResource is a resource wrapper that will return an appropriate resource colorized by the current theme's `DisabledColor` color.

#### func  NewDisabledResource

```go
func NewDisabledResource(res fyne.Resource) *DisabledResource
```
NewDisabledResource creates a resource that adapts to the current theme's DisabledColor setting.

#### func (*DisabledResource) Content

```go
func (res *DisabledResource) Content() []byte
```
Content returns the disabled style content of the correct resource for the current theme

#### func (*DisabledResource) Name

```go
func (res *DisabledResource) Name() string
```
Name returns the resource source name prefixed with `disabled_` (used for caching)

#### func (*DisabledResource) ThemeColorName

```go
func (res *DisabledResource) ThemeColorName() fyne.ThemeColorName
```
ThemeColorName returns the fyne.ThemeColorName that is used as foreground color.


<div class="since">Since: <code>
2.6</code></div>
