---
layout: page
tags: [api]
title: Fyne API "software"
---

# software
---
```go
import "fyne.io/fyne/v2/driver/software"
```


## Usage

#### func  NewCanvas

```go
func NewCanvas() test.WindowlessCanvas
```
NewCanvas creates a new canvas in memory that can render without hardware support

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
