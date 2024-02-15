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

Package canvas contains all of the primitive CanvasObjects that make up a Fyne GUI

The types implemented in this package are used as building blocks in order to build higher order functionality. These types are designed to be non-interactive, by design. If additional functonality is required, it's usually a sign that this type should be used as part of a custom Widget.

## Usage

#### func  Refresh

```go
func Refresh(obj fyne.CanvasObject)
```
Refresh instructs the containing canvas to refresh the specified obj.

#### types

 * [Circle](circle.html)
 * [Image](image.html)
 * [ImageFill](imagefill.html)
 * [Line](line.html)
 * [LinearGradient](lineargradient.html)
 * [RadialGradient](radialgradient.html)
 * [Raster](raster.html)
 * [Rectangle](rectangle.html)
 * [Text](text.html)
