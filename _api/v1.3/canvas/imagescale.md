---
layout: page
tags: [api]
title: Fyne API canvas
---

# canvas
---
```go
import "fyne.io/fyne/canvas"
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
	ImageScaleSmooth ImageScale = 0
	// ImageScalePixels will scale the image using NearestNeighbor filter (or GL equivalent)
	ImageScalePixels ImageScale = 1
)
```
