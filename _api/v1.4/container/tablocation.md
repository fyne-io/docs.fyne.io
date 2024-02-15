---
layout: page
tags: [api]
title: Fyne API "container.TabLocation"
---

# container.TabLocation
---
```go
import "fyne.io/fyne/container"
```

## Usage

#### type TabLocation

```go
type TabLocation = widget.TabLocation
```

TabLocation is the location where the tabs of a tab container should be rendered


<div class="since">Since: <code>
1.4</code></div>

```go
const (
	TabLocationTop TabLocation = iota
	TabLocationLeading
	TabLocationBottom
	TabLocationTrailing
)
```
TabLocation values
