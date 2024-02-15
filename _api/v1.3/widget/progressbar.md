---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type ProgressBar

```go
type ProgressBar struct {
	BaseWidget

	Min, Max, Value float64
}
```

ProgressBar widget creates a horizontal panel that indicates progress

#### func  NewProgressBar

```go
func NewProgressBar() *ProgressBar
```
NewProgressBar creates a new progress bar widget. The default Min is 0 and Max is 1, Values set should be between those numbers. The display will convert this to a percentage.

#### func (*ProgressBar) CreateRenderer

```go
func (p *ProgressBar) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*ProgressBar) MinSize

```go
func (p *ProgressBar) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*ProgressBar) SetValue

```go
func (p *ProgressBar) SetValue(v float64)
```
SetValue changes the current value of this progress bar (from p.Min to p.Max). The widget will be refreshed to indicate the change.
