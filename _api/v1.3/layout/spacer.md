---
layout: page
tags: [api]
title: Fyne API layout
---

# layout
---
```go
import "fyne.io/fyne/layout"
```

## Usage

#### type Spacer

```go
type Spacer struct {
	FixHorizontal bool
	FixVertical   bool
}
```

Spacer is any simple object that can be used in a box layout to space out child objects

#### func (*Spacer) ExpandHorizontal

```go
func (s *Spacer) ExpandHorizontal() bool
```
ExpandHorizontal returns whether or not this spacer expands on the horizontal axis

#### func (*Spacer) ExpandVertical

```go
func (s *Spacer) ExpandVertical() bool
```
ExpandVertical returns whether or not this spacer expands on the vertical axis

#### func (*Spacer) Hide

```go
func (s *Spacer) Hide()
```
Hide removes this Spacer from layout calculations

#### func (*Spacer) MinSize

```go
func (s *Spacer) MinSize() fyne.Size
```
MinSize returns a 0 size as a Spacer can shrink to no actual size

#### func (*Spacer) Move

```go
func (s *Spacer) Move(pos fyne.Position)
```
Move sets a new position for the Spacer - this will be called by the layout

#### func (*Spacer) Position

```go
func (s *Spacer) Position() fyne.Position
```
Position returns the current position of this Spacer

#### func (*Spacer) Refresh

```go
func (s *Spacer) Refresh()
```
Refresh does nothing for a spacer but is part of the CanvasObject definition

#### func (*Spacer) Resize

```go
func (s *Spacer) Resize(size fyne.Size)
```
Resize sets a new size for the Spacer - this will be called by the layout

#### func (*Spacer) Show

```go
func (s *Spacer) Show()
```
Show sets the Spacer to be part of the layout calculations

#### func (*Spacer) Size

```go
func (s *Spacer) Size() fyne.Size
```
Size returns the current size of this Spacer

#### func (*Spacer) Visible

```go
func (s *Spacer) Visible() bool
```
Visible returns true if this spacer should affect the layout
