---
tags: [api]
title: fyne.ThemedResource
slug: themedresource

aliases:
- /api//themedresource
- /api//themedresource.html
- /api/v2.0//themedresource
- /api/v2.0//themedresource.html
- /api/v2.1//themedresource
- /api/v2.1//themedresource.html
- /api/v2.2//themedresource
- /api/v2.2//themedresource.html
- /api/v2.3//themedresource
- /api/v2.3//themedresource.html
- /api/v2.4//themedresource
- /api/v2.4//themedresource.html
- /api/v2.5//themedresource
- /api/v2.5//themedresource.html
- /api/v2.6//themedresource
- /api/v2.6//themedresource.html
- /api/v2.7//themedresource
- /api/v2.7//themedresource.html

package: fyne.io/fyne/v2
---


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

ThemedResource is a version of a resource that can be updated to match a certain theme color. The [ThemeColorName] will be used to look up the color for the current theme and colorize the resource.


<div class="since">Since: <code>
2.5</code></div>
