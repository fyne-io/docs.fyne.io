---
layout: page
tags: [api]
title: Fyne API "layout.SpacerObject"
package: fyne.io/fyne/v2/layout
---

# layout.SpacerObject
---
```go
import "fyne.io/fyne/v2/layout"
```

## Usage

#### type SpacerObject

```go
type SpacerObject interface {
	ExpandVertical() bool
	ExpandHorizontal() bool
}
```

SpacerObject is any object that can be used to space out child objects
