---
tags: [api]
title: widget.Calendar
slug: calendar

aliases:
- /api/widget/calendar
- /api/widget/calendar.html
- /api/v2.0/widget/calendar
- /api/v2.0/widget/calendar.html
- /api/v2.1/widget/calendar
- /api/v2.1/widget/calendar.html
- /api/v2.2/widget/calendar
- /api/v2.2/widget/calendar.html
- /api/v2.3/widget/calendar
- /api/v2.3/widget/calendar.html
- /api/v2.4/widget/calendar
- /api/v2.4/widget/calendar.html
- /api/v2.5/widget/calendar
- /api/v2.5/widget/calendar.html
- /api/v2.6/widget/calendar
- /api/v2.6/widget/calendar.html
- /api/v2.7/widget/calendar
- /api/v2.7/widget/calendar.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Calendar

```go
type Calendar struct {
	BaseWidget

	OnChanged func(time.Time) `json:"-"`
}
```

Calendar creates a new date time picker which returns a time object


<div class="since">Since: <code>
2.6</code></div>

#### func  NewCalendar

```go
func NewCalendar(cT time.Time, changed func(time.Time)) *Calendar
```
NewCalendar creates a calendar instance


<div class="since">Since: <code>
2.6</code></div>

#### func (*Calendar) CreateRenderer

```go
func (c *Calendar) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new WidgetRenderer for this widget. This should not be called by regular code, it is used internally to render a widget.
