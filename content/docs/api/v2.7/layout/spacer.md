---
tags: [api]
title: layout.Spacer
slug: spacer

aliases:
- /api/v2.7/layout/spacer

package: fyne.io/fyne/v2/layout
---


---
```go
import "fyne.io/fyne/v2/layout"
```

#

###

```go
type Spacer struct {
	FixHorizontal bool
	FixVertical   bool
}
```

Spacer is any simple object that can be used in a box layout to space out child objects

###

```go
func (s *Spacer) ExpandHorizontal() bool
```
ExpandHorizontal returns whether or not this spacer expands on the horizontal axis

###

```go
func (s *Spacer) ExpandVertical() bool
```
ExpandVertical returns whether or not this spacer expands on the vertical axis

###

```go
func (s *Spacer) Hide()
```
Hide removes this Spacer from layout calculations

###

```go
func (s *Spacer) MinSize() fyne.Size
```
MinSize returns a 0 size as a Spacer can shrink to no actual size

###

```go
func (s *Spacer) Move(pos fyne.Position)
```
Move sets a new position for the Spacer - this will be called by the layout

###

```go
func (s *Spacer) Position() fyne.Position
```
Position returns the current position of this Spacer

###

```go
func (s *Spacer) Refresh()
```
Refresh does nothing for a spacer but is part of the CanvasObject definition

###

```go
func (s *Spacer) Resize(size fyne.Size)
```
Resize sets a new size for the Spacer - this will be called by the layout

###

```go
func (s *Spacer) Show()
```
Show sets the Spacer to be part of the layout calculations

###

```go
func (s *Spacer) Size() fyne.Size
```
Size returns the current size of this Spacer

###

```go
func (s *Spacer) Visible() bool
```
Visible returns true if this spacer should affect the layout
