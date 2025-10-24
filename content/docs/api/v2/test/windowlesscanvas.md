---
tags: [api]
title: test.WindowlessCanvas
slug: windowlesscanvas

aliases:
- /api/v2.0/test/windowlesscanvas
- /api/v2.1/test/windowlesscanvas
- /api/v2.2/test/windowlesscanvas
- /api/v2.3/test/windowlesscanvas
- /api/v2.4/test/windowlesscanvas
- /api/v2.5/test/windowlesscanvas
- /api/v2.6/test/windowlesscanvas
- /api/v2.7/test/windowlesscanvas

package: fyne.io/fyne/v2/test
---


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
