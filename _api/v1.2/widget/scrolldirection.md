---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
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

```go
const (
	ScrollBoth ScrollDirection = iota
	ScrollHorizontalOnly
	ScrollVerticalOnly
)
```
Constants for valid values of ScrollDirection.
