---
tags: [api]
title: fyne.ScrollDirection
slug: scrolldirection

aliases:
- /api/v2.0//scrolldirection
- /api/v2.1//scrolldirection
- /api/v2.2//scrolldirection
- /api/v2.3//scrolldirection
- /api/v2.4//scrolldirection
- /api/v2.5//scrolldirection
- /api/v2.6//scrolldirection
- /api/v2.7//scrolldirection

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ScrollDirection

```go
type ScrollDirection int
```

ScrollDirection represents the directions in which a scrollable container or widget can scroll its child content.


<div class="since">Since: <code>
2.6</code></div>

```go
const (
	// ScrollBoth supports horizontal and vertical scrolling.
	ScrollBoth ScrollDirection = iota
	// ScrollHorizontalOnly specifies the scrolling should only happen left to right.
	ScrollHorizontalOnly
	// ScrollVerticalOnly specifies the scrolling should only happen top to bottom.
	ScrollVerticalOnly
	// ScrollNone turns off scrolling for this container.
	ScrollNone
)
```
Constants for valid values of ScrollDirection used in containers and widgets.
