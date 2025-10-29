---
tags: [api]
title: software (package)
slug: (package)

aliases:
- /api/driver/software
- /api/driver/software.html
- /api/v2.0/driver/software
- /api/v2.0/driver/software.html
- /api/v2.1/driver/software
- /api/v2.1/driver/software.html
- /api/v2.2/driver/software
- /api/v2.2/driver/software.html
- /api/v2.3/driver/software
- /api/v2.3/driver/software.html
- /api/v2.4/driver/software
- /api/v2.4/driver/software.html
- /api/v2.5/driver/software
- /api/v2.5/driver/software.html
- /api/v2.6/driver/software
- /api/v2.6/driver/software.html
- /api/v2.7/driver/software
- /api/v2.7/driver/software.html

package: fyne.io/fyne/v2/driver/software
---


---
```go
import "fyne.io/fyne/v2/driver/software"
```


## Usage

#### func  NewCanvas

```go
func NewCanvas() test.WindowlessCanvas
```
NewCanvas creates a new canvas in memory that can render without hardware support.

#### func  NewTransparentCanvas

```go
func NewTransparentCanvas() test.WindowlessCanvas
```
NewTransparentCanvas creates a new canvas in memory that can render without hardware support without a background color.


<div class="since">Since: <code>
2.2</code></div>

#### func  Render

```go
func Render(obj fyne.CanvasObject, t fyne.Theme) image.Image
```
Render takes a canvas object and renders it to a regular Go image using the provided Theme. The returned image will be set to the object's minimum size. Use the theme.LightTheme() or theme.DarkTheme() to access the builtin themes.

#### func  RenderCanvas

```go
func RenderCanvas(c fyne.Canvas, t fyne.Theme) image.Image
```
RenderCanvas takes a canvas and renders it to a regular Go image using the provided Theme. This is the same as setting the application theme and then calling Canvas.Capture().

#### types
