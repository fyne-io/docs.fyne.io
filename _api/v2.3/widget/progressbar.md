---
layout: page
tags: [api]
title: Fyne API "widget.ProgressBar"
---

# widget.ProgressBar
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ProgressBar

```go
type ProgressBar struct {
	BaseWidget

	Min, Max, Value float64

	// TextFormatter can be used to have a custom format of progress text.
	// If set, it overrides the percentage readout and runs each time the value updates.
	//
	// Since: 1.4
	TextFormatter func() string `json:"-"`
}
```

ProgressBar widget creates a horizontal panel that indicates progress

#### func  NewProgressBar

```go
func NewProgressBar() *ProgressBar
```
NewProgressBar creates a new progress bar widget. The default Min is 0 and Max is 1, Values set should be between those numbers. The display will convert this to a percentage.

#### func  NewProgressBarWithData

```go
func NewProgressBarWithData(data binding.Float) *ProgressBar
```
NewProgressBarWithData returns a progress bar connected with the specified data source.


<div class="since">Since: <code>
2.0</code></div>

#### func (*ProgressBar) Bind

```go
func (p *ProgressBar) Bind(data binding.Float)
```
Bind connects the specified data source to this ProgressBar. The current value will be displayed and any changes in the data will cause the widget to update.


<div class="since">Since: <code>
2.0</code></div>

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

#### func (*ProgressBar) Unbind

```go
func (p *ProgressBar) Unbind()
```
Unbind disconnects any configured data source from this ProgressBar. The current value will remain at the last value of the data source.


<div class="since">Since: <code>
2.0</code></div>
