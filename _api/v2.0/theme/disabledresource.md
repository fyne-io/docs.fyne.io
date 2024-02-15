---
layout: page
tags: [api]
title: Fyne API "theme.DisabledResource"
---

# theme.DisabledResource
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

DisabledResource is a resource wrapper that will return an appropriate resource colorized by the current theme's `DisabledIconColor` color.

#### func  NewDisabledResource

```go
func NewDisabledResource(res fyne.Resource) *DisabledResource
```
NewDisabledResource creates a resource that adapts to the current theme's DisabledIconColor setting.

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
