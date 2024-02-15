---
layout: page
tags: [api]
title: Fyne API "container.AppTabs"
---

# container.AppTabs
---
```go
import "fyne.io/fyne/container"
```

## Usage

#### type AppTabs

```go
type AppTabs = widget.TabContainer
```

AppTabs container is used to split your application into various different areas identified by tabs. The tabs contain text and/or an icon and allow the user to switch between the content specified in each TabItem. Each item is represented by a button at the edge of the container.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewAppTabs

```go
func NewAppTabs(items ...*TabItem) *AppTabs
```
NewAppTabs creates a new tab container that allows the user to choose between different areas of an app.


<div class="since">Since: <code>
1.4</code></div>
