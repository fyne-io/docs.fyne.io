---
layout: page
tags: [api]
title: Fyne API "widget.Slider"
---

# widget.Slider
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Slider

```go
type Slider struct {
	BaseWidget

	Value float64
	Min   float64
	Max   float64
	Step  float64

	Orientation Orientation
	OnChanged   func(float64)
}
```

Slider is a widget that can slide between two fixed values.

#### func  NewSlider

```go
func NewSlider(min, max float64) *Slider
```
NewSlider returns a basic slider.

#### func  NewSliderWithData

```go
func NewSliderWithData(min, max float64, data binding.Float) *Slider
```
NewSliderWithData returns a slider connected with the specified data source.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Slider) Bind

```go
func (s *Slider) Bind(data binding.Float)
```
Bind connects the specified data source to this Slider. The current value will be displayed and any changes in the data will cause the widget to update. User interactions with this Slider will set the value into the data source.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Slider) CreateRenderer

```go
func (s *Slider) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer links this widget to its renderer.

#### func (*Slider) DragEnd

```go
func (s *Slider) DragEnd()
```
DragEnd function.

#### func (*Slider) Dragged

```go
func (s *Slider) Dragged(e *fyne.DragEvent)
```
Dragged function.

#### func (*Slider) MinSize

```go
func (s *Slider) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Slider) SetValue

```go
func (s *Slider) SetValue(value float64)
```
SetValue updates the value of the slider and clamps the value to be within the range.

#### func (*Slider) Unbind

```go
func (s *Slider) Unbind()
```
Unbind disconnects any configured data source from this Slider. The current value will remain at the last value of the data source.


<div class="since">Since: <code>
2.0</code></div>
