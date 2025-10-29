---
tags: [api]
title: layout.SpacerObject
slug: spacerobject

aliases:
- /api/layout/spacerobject
- /api/layout/spacerobject.html
- /api/v2.0/layout/spacerobject
- /api/v2.0/layout/spacerobject.html
- /api/v2.1/layout/spacerobject
- /api/v2.1/layout/spacerobject.html
- /api/v2.2/layout/spacerobject
- /api/v2.2/layout/spacerobject.html
- /api/v2.3/layout/spacerobject
- /api/v2.3/layout/spacerobject.html
- /api/v2.4/layout/spacerobject
- /api/v2.4/layout/spacerobject.html
- /api/v2.5/layout/spacerobject
- /api/v2.5/layout/spacerobject.html
- /api/v2.6/layout/spacerobject
- /api/v2.6/layout/spacerobject.html
- /api/v2.7/layout/spacerobject
- /api/v2.7/layout/spacerobject.html

package: fyne.io/fyne/v2/layout
---


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
