---
tags: [api]
title: widget.Calendar
slug: calendar

aliases:
- /api/v2.7/widget/calendar

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type Calendar struct {
	BaseWidget

	OnChanged func(time.Time) `json:"-"`
}
```

Calendar creates a new date time picker which returns a time object


<div class="since">Since: <code>
2.6</code></div>

###

```go
func NewCalendar(cT time.Time, changed func(time.Time)) *Calendar
```
NewCalendar creates a calendar instance


<div class="since">Since: <code>
2.6</code></div>

###

```go
func (c *Calendar) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new WidgetRenderer for this widget. This should not be called by regular code, it is used internally to render a widget.
