---
tags: [api]
title: fyne.ThemedResource
slug: themedresource

aliases:
- /api/v2.7//themedresource

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type ThemedResource interface {
	Resource
	ThemeColorName() ThemeColorName
}
```

ThemedResource is a version of a resource that can be updated to match a certain theme color. The [ThemeColorName] will be used to look up the color for the current theme and colorize the resource.


<div class="since">Since: <code>
2.5</code></div>
