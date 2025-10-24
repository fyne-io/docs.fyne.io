---
tags: [api]
title: test.SoftwarePainter
slug: softwarepainter

aliases:
- /api/v2.7/test/softwarepainter

package: fyne.io/fyne/v2/test
---


---
```go
import "fyne.io/fyne/v2/test"
```

#

###

```go
type SoftwarePainter interface {
	Paint(fyne.Canvas) image.Image
}
```

SoftwarePainter describes a simple type that can render canvases
