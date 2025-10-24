---
tags: [api]
title: widget.DateEntry
slug: dateentry

aliases:
- /api/v2.0/widget/dateentry
- /api/v2.1/widget/dateentry
- /api/v2.2/widget/dateentry
- /api/v2.3/widget/dateentry
- /api/v2.4/widget/dateentry
- /api/v2.5/widget/dateentry
- /api/v2.6/widget/dateentry
- /api/v2.7/widget/dateentry

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type DateEntry

```go
type DateEntry struct {
	Entry
	Date      *time.Time
	OnChanged func(*time.Time) `json:"-"`
}
```

DateEntry is an input field which supports selecting from a fixed set of options.


<div class="since">Since: <code>
2.6</code></div>

#### func  NewDateEntry

```go
func NewDateEntry() *DateEntry
```
NewDateEntry creates a date input where the date can be selected or typed.


<div class="since">Since: <code>
2.6</code></div>

#### func (*DateEntry) CreateRenderer

```go
func (e *DateEntry) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new renderer for this select entry.

#### func (*DateEntry) Disable

```go
func (e *DateEntry) Disable()
```
Disable this widget so that it cannot be interacted with, updating any style appropriately.

#### func (*DateEntry) Enable

```go
func (e *DateEntry) Enable()
```
Enable this widget, updating any style or features appropriately.

#### func (*DateEntry) MinSize

```go
func (e *DateEntry) MinSize() fyne.Size
```
MinSize returns the minimal size of the select entry.

#### func (*DateEntry) Move

```go
func (e *DateEntry) Move(pos fyne.Position)
```
Move changes the relative position of the date entry.

#### func (*DateEntry) Resize

```go
func (e *DateEntry) Resize(size fyne.Size)
```
Resize changes the size of the date entry.

#### func (*DateEntry) SetDate

```go
func (e *DateEntry) SetDate(d *time.Time)
```
SetDate will update the widget to a specific date. You can pass nil to unselect a date.
