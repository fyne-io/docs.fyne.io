---
layout: page
tags: [api]
title: Fyne API "container.ScrollDirection"
---

# container.ScrollDirection
---
```go
import "fyne.io/fyne/container"
```

## Usage

#### type ScrollDirection

```go
type ScrollDirection = widget.ScrollDirection
```

ScrollDirection represents the directions in which a Scroll container can scroll its child content.


<div class="since">Since: <code>
1.4</code></div>

```go
const (
	ScrollBoth ScrollDirection = iota
	ScrollHorizontalOnly
	ScrollVerticalOnly
)
```
Constants for valid values of ScrollDirection.
