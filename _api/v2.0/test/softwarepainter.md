---
layout: page
tags: [api]
title: Fyne API "test.SoftwarePainter"
---

# test.SoftwarePainter
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
