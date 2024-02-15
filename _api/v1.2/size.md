---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
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

#### func  NewSize

```go
func NewSize(w int, h int) Size
```
NewSize returns a newly allocated Size of the specified dimensions.

#### func (Size) Add

```go
func (s1 Size) Add(s2 Size) Size
```
Add returns a new Size that is the result of increasing the current size by s2 Width and Height.

#### func (Size) Subtract

```go
func (s1 Size) Subtract(s2 Size) Size
```
Subtract returns a new Size that is the result of decreasing the current size by s2 Width and Height.

#### func (Size) Union

```go
func (s1 Size) Union(s2 Size) Size
```
Union returns a new Size that is the maximum of the current Size and s2.
