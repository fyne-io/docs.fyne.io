---
tags: [api]
title: canvas.ImageScale
slug: imagescale

aliases:
- /api/v2/canvas/imagescale.html
- /api/v2.0/canvas/imagescale
- /api/v2.0/canvas/imagescale.html
- /api/v2.1/canvas/imagescale
- /api/v2.1/canvas/imagescale.html
- /api/v2.2/canvas/imagescale
- /api/v2.2/canvas/imagescale.html
- /api/v2.3/canvas/imagescale
- /api/v2.3/canvas/imagescale.html
- /api/v2.4/canvas/imagescale
- /api/v2.4/canvas/imagescale.html
- /api/v2.5/canvas/imagescale
- /api/v2.5/canvas/imagescale.html
- /api/v2.6/canvas/imagescale
- /api/v2.6/canvas/imagescale.html
- /api/v2.7/canvas/imagescale
- /api/v2.7/canvas/imagescale.html

package: fyne.io/fyne/v2/canvas
---


---
```go
import "fyne.io/fyne/v2/canvas"
```

## Usage

#### type ImageScale

```go
type ImageScale int32
```

ImageScale defines the different scaling filters used to scaling images

```go
const (
	// ImageScaleSmooth will scale the image using ApproxBiLinear filter (or GL equivalent)
	ImageScaleSmooth ImageScale = iota
	// ImageScalePixels will scale the image using NearestNeighbor filter (or GL equivalent)
	ImageScalePixels
	// ImageScaleFastest will scale the image using hardware GPU if available
	//
	// Since: 2.0
	ImageScaleFastest
)
```
