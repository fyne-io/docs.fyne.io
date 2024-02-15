---
layout: page
tags: [api]
title: Fyne API "fyne.Size"
---

# fyne.Size
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Size

```go
type Size struct {
	Width  int // The number of units along the X axis.
	Height int // The number of units along the Y axis.
}
```

Size describes something with width and height.

#### func  MeasureText

```go
func MeasureText(text string, size int, style TextStyle) Size
```
MeasureText uses the current driver to calculate the size of text when rendered.

#### func  NewSize

```go
func NewSize(w int, h int) Size
```
NewSize returns a newly allocated Size of the specified dimensions.

#### func (Size) Add

```go
func (s Size) Add(s2 Size) Size
```
Add returns a new Size that is the result of increasing the current size by s2 Width and Height.

#### func (Size) IsZero

```go
func (s Size) IsZero() bool
```
IsZero returns whether the Size has zero width and zero height.

#### func (Size) Max

```go
func (s Size) Max(s2 Size) Size
```
Max returns a new Size that is the maximum of the current Size and s2.

#### func (Size) Min

```go
func (s Size) Min(s2 Size) Size
```
Min returns a new Size that is the minimum of the current Size and s2.

#### func (Size) Subtract

```go
func (s Size) Subtract(s2 Size) Size
```
Subtract returns a new Size that is the result of decreasing the current size by s2 Width and Height.

#### func (Size) Union

```go
func (s Size) Union(s2 Size) Size
```
Union returns a new Size that is the maximum of the current Size and s2.


<div class="deprecated">
Deprecated: use Max() instead</div>
