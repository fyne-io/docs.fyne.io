---
tags: [api]
title: fyne.Size
slug: size

aliases:
- /api/v2/fyne/size.html
- /api/v2.0//size
- /api/v2.0//size.html
- /api/v2.1//size
- /api/v2.1//size.html
- /api/v2.2//size
- /api/v2.2//size.html
- /api/v2.3//size
- /api/v2.3//size.html
- /api/v2.4//size
- /api/v2.4//size.html
- /api/v2.5//size
- /api/v2.5//size.html
- /api/v2.6//size
- /api/v2.6//size.html
- /api/v2.7//size
- /api/v2.7//size.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Size

```go
type Size struct {
	Width  float32 // The number of units along the X axis.
	Height float32 // The number of units along the Y axis.
}
```

Size describes something with width and height.

#### func  MeasureText

```go
func MeasureText(text string, size float32, style TextStyle) Size
```
MeasureText uses the current driver to calculate the size of text when rendered. The font used will be read from the current app's theme.

#### func  NewSize

```go
func NewSize(w float32, h float32) Size
```
NewSize returns a newly allocated Size of the specified dimensions.

#### func  NewSquareSize

```go
func NewSquareSize(side float32) Size
```
NewSquareSize returns a newly allocated Size with the same width and height.


<div class="since">Since: <code>
2.4</code></div>

#### func (Size) Add

```go
func (s Size) Add(v Vector2) Size
```
Add returns a new Size that is the result of increasing the current size by s2 Width and Height.

#### func (Size) AddWidthHeight

```go
func (s Size) AddWidthHeight(width, height float32) Size
```
AddWidthHeight returns a new Size by adding width and height to the current one.

#### func (Size) Components

```go
func (s Size) Components() (float32, float32)
```
Components returns the Width and Height elements of this Size

#### func (Size) IsZero

```go
func (s Size) IsZero() bool
```
IsZero returns whether the Size has zero width and zero height.

#### func (Size) Max

```go
func (s Size) Max(v Vector2) Size
```
Max returns a new [Size] that is the maximum of the current Size and s2.

#### func (Size) Min

```go
func (s Size) Min(v Vector2) Size
```
Min returns a new [Size] that is the minimum of s and v.

#### func (Size) Subtract

```go
func (s Size) Subtract(v Vector2) Size
```
Subtract returns a new Size that is the result of decreasing the current size by s2 Width and Height.

#### func (Size) SubtractWidthHeight

```go
func (s Size) SubtractWidthHeight(width, height float32) Size
```
SubtractWidthHeight returns a new Size by subtracting width and height from the current one.
