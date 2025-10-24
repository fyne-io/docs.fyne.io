---
tags: [api]
title: layout.SpacerObject
slug: spacerobject

aliases:
- /api/v2.7/layout/spacerobject

package: fyne.io/fyne/v2/layout
---


---
```go
import "fyne.io/fyne/v2/layout"
```

#

###

```go
type SpacerObject interface {
	ExpandVertical() bool
	ExpandHorizontal() bool
}
```

SpacerObject is any object that can be used to space out child objects
