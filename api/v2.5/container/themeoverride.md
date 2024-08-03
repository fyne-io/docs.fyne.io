---
layout: page
tags: [api]
title: Fyne API "container.ThemeOverride"
package: fyne.io/fyne/v2/container
---

# container.ThemeOverride
---
```go
import "fyne.io/fyne/v2/container"
```

## Usage

#### type ThemeOverride

```go
type ThemeOverride struct {
	widget.BaseWidget

	Content fyne.CanvasObject
	Theme   fyne.Theme
}
```

ThemeOverride is a container where the child widgets are themed by the specified theme. Containers will be traversed and all child widgets will reflect the theme in this container. This should be used sparingly to avoid a jarring user experience.


<div class="since">Since: <code>
2.5</code></div>

#### func  NewThemeOverride

```go
func NewThemeOverride(obj fyne.CanvasObject, th fyne.Theme) *ThemeOverride
```
NewThemeOverride provides a container where the child widgets are themed by the specified theme. Containers will be traversed and all child widgets will reflect the theme in this container. This should be used sparingly to avoid a jarring user experience.

If the content `obj` of this theme override is a container and items are later added to the container or any sub-containers ensure that you call `Refresh()` on this `ThemeOverride` to ensure the new items match the theme.


<div class="since">Since: <code>
2.5</code></div>

#### func (*ThemeOverride) CreateRenderer

```go
func (t *ThemeOverride) CreateRenderer() fyne.WidgetRenderer
```

#### func (*ThemeOverride) Refresh

```go
func (t *ThemeOverride) Refresh()
```
