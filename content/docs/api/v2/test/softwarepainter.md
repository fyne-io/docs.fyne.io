---
tags: [api]
title: test.SoftwarePainter
slug: softwarepainter

aliases:
- /api/v2.0/test/softwarepainter
- /api/v2.1/test/softwarepainter
- /api/v2.2/test/softwarepainter
- /api/v2.3/test/softwarepainter
- /api/v2.4/test/softwarepainter
- /api/v2.5/test/softwarepainter
- /api/v2.6/test/softwarepainter
- /api/v2.7/test/softwarepainter

package: fyne.io/fyne/v2/test
---


---
```go
import "fyne.io/fyne/v2/test"
```

## Usage

#### type SoftwarePainter

```go
type SoftwarePainter interface {
	Paint(fyne.Canvas) image.Image
}
```

SoftwarePainter describes a simple type that can render canvases
