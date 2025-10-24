---
tags: [api]
title: canvas (package)
slug: (package)

aliases:
- /api/v2.0/canvas
- /api/v2.1/canvas
- /api/v2.2/canvas
- /api/v2.3/canvas
- /api/v2.4/canvas
- /api/v2.5/canvas
- /api/v2.6/canvas
- /api/v2.7/canvas

package: fyne.io/fyne/v2/canvas
---


---
```go
import "fyne.io/fyne/v2/canvas"
```

Package canvas contains all of the primitive CanvasObjects that make up a Fyne GUI.

The types implemented in this package are used as building blocks in order to build higher order functionality. These types are designed to be non-interactive, by design. If additional functionality is required, it's usually a sign that this type should be used as part of a custom widget.

## Usage

```go
const (
	// DurationStandard is the time a standard interface animation will run.
	//
	// Since: 2.0
	DurationStandard = time.Millisecond * 300
	// DurationShort is the time a subtle or small transition should use.
	//
	// Since: 2.0
	DurationShort = time.Millisecond * 150
)
```

```go
const (
	// RadiusMaximum can be applied to a canvas corner radius to achieve fully rounded corners.
	// This constant represents the maximum possible corner radius, resulting in a circular appearance.
	// Since: 2.7
	RadiusMaximum float32 = math.MaxFloat32
)
```

#### func  NewColorRGBAAnimation

```go
func NewColorRGBAAnimation(start, stop color.Color, d time.Duration, fn func(color.Color)) *fyne.Animation
```
NewColorRGBAAnimation sets up a new animation that will transition from the start to stop Color over the specified Duration. The colour transition will move linearly through the RGB colour space. The content of fn should apply the color values to an object and refresh it. You should call Start() on the returned animation to start it.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewPositionAnimation

```go
func NewPositionAnimation(start, stop fyne.Position, d time.Duration, fn func(fyne.Position)) *fyne.Animation
```
NewPositionAnimation sets up a new animation that will transition from the start to stop Position over the specified Duration. The content of fn should apply the position value to an object for the change to be visible. You should call Start() on the returned animation to start it.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewSizeAnimation

```go
func NewSizeAnimation(start, stop fyne.Size, d time.Duration, fn func(fyne.Size)) *fyne.Animation
```
NewSizeAnimation sets up a new animation that will transition from the start to stop Size over the specified Duration. The content of fn should apply the size value to an object for the change to be visible. You should call Start() on the returned animation to start it.


<div class="since">Since: <code>
2.0</code></div>

#### func  RecolorSVG

```go
func RecolorSVG(svgContent []byte, color color.Color) ([]byte, error)
```
RecolorSVG takes a []byte containing SVG content, and returns new SVG content, re-colorized to be monochrome with the given color. The content can be assigned to a new fyne.StaticResource with an appropriate name to be used in a widget.Button, canvas.Image, etc.

If an error occurs, the returned content will be the original un-modified content, and a non-nil error is returned.


<div class="since">Since: <code>
2.6</code></div>

#### func  Refresh

```go
func Refresh(obj fyne.CanvasObject)
```
Refresh instructs the containing canvas to refresh the specified obj.

#### types

 * [Arc](arc.html)
 * [Circle](circle.html)
 * [Image](image.html)
 * [ImageFill](imagefill.html)
 * [ImageScale](imagescale.html)
 * [Line](line.html)
 * [LinearGradient](lineargradient.html)
 * [Polygon](polygon.html)
 * [RadialGradient](radialgradient.html)
 * [Raster](raster.html)
 * [Rectangle](rectangle.html)
 * [Text](text.html)
