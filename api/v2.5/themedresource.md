---
layout: page
tags: [api]
title: Fyne API "fyne.ThemedResource"
package: fyne.io/fyne/v2
---

# fyne.ThemedResource
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ThemedResource

```go
type ThemedResource interface {
	Resource
	ThemeColorName() ThemeColorName
}
```

ThemedResource is a version of a resource that can be updated to match a certain theme colour. The `ThemeColorName` will be used to look up the color for the current theme and colorize the resource.


<div class="since">Since: <code>
2.5</code></div>
