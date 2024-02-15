---
layout: page
tags: [api]
title: Fyne API "widget.ProgressBarInfinite"
---

# widget.ProgressBarInfinite
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ProgressBarInfinite

```go
type ProgressBarInfinite struct {
	BaseWidget
}
```

ProgressBarInfinite widget creates a horizontal panel that indicates waiting indefinitely An infinite progress bar loops 0% -> 100% repeatedly until Stop() is called

#### func  NewProgressBarInfinite

```go
func NewProgressBarInfinite() *ProgressBarInfinite
```
NewProgressBarInfinite creates a new progress bar widget that loops indefinitely from 0% -> 100% SetValue() is not defined for infinite progress bar To stop the looping progress and set the progress bar to 100%, call ProgressBarInfinite.Stop()

#### func (*ProgressBarInfinite) CreateRenderer

```go
func (p *ProgressBarInfinite) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*ProgressBarInfinite) Hide

```go
func (p *ProgressBarInfinite) Hide()
```
Hide this widget, if it was previously visible

#### func (*ProgressBarInfinite) MinSize

```go
func (p *ProgressBarInfinite) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*ProgressBarInfinite) Running

```go
func (p *ProgressBarInfinite) Running() bool
```
Running returns the current state of the infinite progress animation

#### func (*ProgressBarInfinite) Show

```go
func (p *ProgressBarInfinite) Show()
```
Show this widget, if it was previously hidden

#### func (*ProgressBarInfinite) Start

```go
func (p *ProgressBarInfinite) Start()
```
Start the infinite progress bar animation

#### func (*ProgressBarInfinite) Stop

```go
func (p *ProgressBarInfinite) Stop()
```
Stop the infinite progress bar animation
