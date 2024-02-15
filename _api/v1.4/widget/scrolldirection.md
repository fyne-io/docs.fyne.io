---
layout: page
tags: [api]
title: Fyne API "widget.ScrollDirection"
---

# widget.ScrollDirection
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type ScrollDirection

```go
type ScrollDirection int
```

ScrollDirection represents the directions in which a ScrollContainer can scroll its child content.


<div class="deprecated">
Deprecated: use container.ScrollDirection instead.</div>

```go
const (
	// Deprecated: use container.ScrollBoth instead
	ScrollBoth ScrollDirection = iota
	// Deprecated: use container.ScrollHorizontalOnly instead
	ScrollHorizontalOnly
	// Deprecated: use container.ScrollVerticalOnly instead
	ScrollVerticalOnly
)
```
Constants for valid values of ScrollDirection.
