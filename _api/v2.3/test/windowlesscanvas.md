---
layout: page
tags: [api]
title: Fyne API "test.WindowlessCanvas"
---

# test.WindowlessCanvas
---
```go
import "fyne.io/fyne/v2/test"
```

## Usage

#### type WindowlessCanvas

```go
type WindowlessCanvas interface {
	fyne.Canvas

	Padded() bool
	Resize(fyne.Size)
	SetPadded(bool)
	SetScale(float32)
}
```

WindowlessCanvas provides functionality for a canvas to operate without a window

#### func  NewCanvas

```go
func NewCanvas() WindowlessCanvas
```
NewCanvas returns a single use in-memory canvas used for testing. This canvas has no painter so calls to Capture() will return a blank image.

#### func  NewCanvasWithPainter

```go
func NewCanvasWithPainter(painter SoftwarePainter) WindowlessCanvas
```
NewCanvasWithPainter allows creation of an in-memory canvas with a specific painter. The painter will be used to render in the Capture() call.

#### func  NewTransparentCanvasWithPainter

```go
func NewTransparentCanvasWithPainter(painter SoftwarePainter) WindowlessCanvas
```
NewTransparentCanvasWithPainter allows creation of an in-memory canvas with a specific painter without a background color. The painter will be used to render in the Capture() call.


<div class="since">Since: <code>
2.2</code></div>
